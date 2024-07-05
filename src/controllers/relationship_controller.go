package controllers

import (
	"ca-server/src/database"
	"ca-server/src/enums"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
	"time"
)

type UpdateRelationship struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}

// func create friend request
func CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	current_user_id := r.Context().Value("id").(uint)

	// get user id from request body
	var user_id uint
	err := json.NewDecoder(r.Body).Decode(&user_id)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	var relationship models.Relationship
	relationship.SenderID = current_user_id
	relationship.ReceiverID = user_id
	relationship.CreatedAt = time.Now()
	relationship.Status = string(enums.FRIEND_REQUEST_PENDING)

	// save friend request
	if err := database.DB.Save(&relationship).Error; err != nil {
		http.Error(w, "Failed to update relationship information", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func accept friend request
func AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {

	// get request id from request body
	var request_id uint
	err := json.NewDecoder(r.Body).Decode(&request_id)
	if err != nil {
		http.Error(w, "Failed to decode  id", http.StatusBadRequest)
		return
	}

	// get request relationship from database
	var relationship models.Relationship
	if err := database.DB.Where("id = ?", request_id).First(&relationship).Error; err != nil {
		http.Error(w, "Relationship not found", http.StatusUnauthorized)
		return
	}

	relationship.Status = string(enums.FRIEND)

	w.WriteHeader(http.StatusOK)
}
