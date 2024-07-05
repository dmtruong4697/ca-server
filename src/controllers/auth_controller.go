package controllers

import (
	"ca-server/src/database"
	"ca-server/src/enums"
	"ca-server/src/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
)

type ValidateEmailRequestBody struct {
	Email        string `json:"email"`
	ValidateCode string `json:"validate_code"`
}

type LoginRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DeviceToken string `json:"device_token"`
}

type LogoutRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtKey = []byte("20204697")

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// generate validate code
func generateRandomCode(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// send email func
func SendEmail(to, subject, body string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", "duongminhtruong2002.lequydon@gmail.com")

	mailer.SetHeader("To", to)

	mailer.SetHeader("Subject", subject)

	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer("smtp.example.com", 587, "duongminhtruong2002.lequydon@gmail.com", "jhda naqz lyrp eozp")

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingUser := models.User{}
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusBadRequest)
		return
	}

	validateCode := generateRandomCode(6)

	user.ValidateCode = validateCode
	user.AccountStatus = string(enums.PENDING)

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send email with validate code
	header := "Validate Your Email"
	body := "Validate code:" + validateCode
	SendEmail(user.Email, header, body)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// func validate email
func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var validateEmailRequestBody ValidateEmailRequestBody
	if err := json.NewDecoder(r.Body).Decode(&validateEmailRequestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := database.DB.Where("email = ?", validateEmailRequestBody.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if dbUser.ValidateCode == validateEmailRequestBody.ValidateCode {
		dbUser.AccountStatus = string(enums.VERIFIED)
		dbUser.ValidateCode = ""

		if err := database.DB.Save(&dbUser).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := "Email validation successful. Your account has been validated."
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	} else {
		http.Error(w, "Invalid validation code", http.StatusBadRequest)
	}
}

// func login
func Login(w http.ResponseWriter, r *http.Request) {
	var userRequest LoginRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := database.DB.Where("email = ?", userRequest.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if dbUser.Password != userRequest.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// set device token
	dbUser.DeviceToken = userRequest.DeviceToken
	if err := database.DB.Save(&dbUser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &LoginClaims{
		ID:    dbUser.ID,
		Email: dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(dbUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseData := map[string]interface{}{
		"token": tokenString,
		"user":  string(jsonUser),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func logout
func Logout(w http.ResponseWriter, r *http.Request) {
	var userRequest LogoutRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser models.User
	if err := database.DB.Where("email = ?", userRequest.Email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if dbUser.Password != userRequest.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// set device token
	dbUser.DeviceToken = ""
	if err := database.DB.Save(&dbUser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := "Logout successful."
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
