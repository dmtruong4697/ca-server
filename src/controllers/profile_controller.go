package controllers

import (
	"ca-server/src/database"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
	"time"
)

type UpdatedProfile struct {
	UserName    string    `json:"user_name"`
	PhoneNumber string    `json:"phone_number"`
	AvatarImage string    `json:"avatar_image"`
	HashtagName string    `json:"hashtag_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UpdatedPasword struct {
	Password string `json:"password"`
}

// func get profile infomation
func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("id").(string)

	var dbUser models.User
	if err := database.DB.Where("id = ?", user_id).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbUser); err != nil {
		http.Error(w, "Failed to encode user info", http.StatusInternalServerError)
	}
}

// func update profile infomation
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var dbUser models.User
	if err := database.DB.Where("email = ?", email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var updatedProfile UpdatedProfile
	err := json.NewDecoder(r.Body).Decode(&updatedProfile)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	// update user profile
	dbUser.UserName = updatedProfile.UserName
	dbUser.AvatarImage = updatedProfile.AvatarImage
	dbUser.Gender = updatedProfile.Gender
	dbUser.DateOfBirth = updatedProfile.DateOfBirth
	dbUser.PhoneNumber = updatedProfile.PhoneNumber
	dbUser.HashtagName = updatedProfile.HashtagName

	// save update
	if err := database.DB.Save(&dbUser).Error; err != nil {
		http.Error(w, "Failed to update user information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbUser); err != nil {
		http.Error(w, "Failed to encode user info", http.StatusInternalServerError)
	}
}

// func update password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var dbUser models.User
	if err := database.DB.Where("email = ?", email).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var updatedPassword UpdatedPasword
	err := json.NewDecoder(r.Body).Decode(&updatedPassword)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	// update account password
	dbUser.Password = updatedPassword.Password

	// save update
	if err := database.DB.Save(&dbUser).Error; err != nil {
		http.Error(w, "Failed to update user information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbUser); err != nil {
		http.Error(w, "Failed to encode user info", http.StatusInternalServerError)
	}
}
