package controllers

import (
	"ca-server/src/database"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
	"time"
)

type GetUserInfoRequestBody struct {
	ID int `json:"id"`
}

type GetUserInfoResponce struct {
	ID            int                 `json:"id"`
	UserName      string              `json:"user_name"`
	Email         string              `json:"email"`
	PhoneNumber   string              `json:"phone_number"`
	AvatarImage   string              `json:"avatar_image"`
	HashtagName   string              `json:"hashtag_name"`
	CreatedAt     time.Time           `json:"created_at"`
	AccountStatus string              `json:"account_status"`
	LastActive    time.Time           `json:"last_active"`
	Gender        string              `json:"gender"`
	DateOfBirth   time.Time           `json:"date_of_birth"`
	Relationship  models.Relationship `json:"relationship"`
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	current_user_id := r.Context().Value("id").(int)

	// get user id from request body
	var user_id uint
	err := json.NewDecoder(r.Body).Decode(&user_id)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	// get user from database
	var dbUser models.User
	if err := database.DB.Where("id = ?", user_id).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	var relationship models.Relationship
	if err := database.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", current_user_id, user_id, user_id, current_user_id).First(&relationship).Error; err != nil {

	}

	var responce GetUserInfoResponce
	responce.ID = int(dbUser.ID)
	responce.AccountStatus = dbUser.AccountStatus
	responce.AvatarImage = dbUser.AvatarImage
	responce.CreatedAt = dbUser.CreatedAt
	responce.DateOfBirth = dbUser.DateOfBirth
	responce.Email = dbUser.Email
	responce.Gender = dbUser.Gender
	responce.HashtagName = dbUser.HashtagName
	responce.UserName = dbUser.UserName
	responce.PhoneNumber = dbUser.PhoneNumber
	responce.LastActive = dbUser.LastActive
	responce.Relationship = relationship

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responce); err != nil {
		http.Error(w, "Failed to encode responce info", http.StatusInternalServerError)
	}
}

func GetUserInfoFunction(current_user_id uint, user_id uint) GetUserInfoResponce {

	// get user from database
	var dbUser models.User
	if err := database.DB.Where("id = ?", user_id).First(&dbUser).Error; err != nil {

	}

	var relationship models.Relationship
	if err := database.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", current_user_id, user_id, user_id, current_user_id).First(&relationship).Error; err != nil {

	}

	var responce GetUserInfoResponce
	responce.ID = int(dbUser.ID)
	responce.AccountStatus = dbUser.AccountStatus
	responce.AvatarImage = dbUser.AvatarImage
	responce.CreatedAt = dbUser.CreatedAt
	responce.DateOfBirth = dbUser.DateOfBirth
	responce.Email = dbUser.Email
	responce.Gender = dbUser.Gender
	responce.HashtagName = dbUser.HashtagName
	responce.UserName = dbUser.UserName
	responce.PhoneNumber = dbUser.PhoneNumber
	responce.LastActive = dbUser.LastActive
	responce.Relationship = relationship

	return responce
}
