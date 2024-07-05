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
