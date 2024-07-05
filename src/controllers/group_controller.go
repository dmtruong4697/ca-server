package controllers

import (
	"ca-server/src/database"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
)

// create personal group
func CreatePersonalGroup(w http.ResponseWriter, r *http.Request) {
	current_user_id := r.Context().Value("id").(uint)

	// get user id from request body
	var user_id uint
	err := json.NewDecoder(r.Body).Decode(&user_id)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	// find relationship
	var relationship models.Relationship
	if err := database.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", current_user_id, user_id, user_id, current_user_id).First(&relationship).Error; err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(relationship); err != nil {
		http.Error(w, "Failed to encode relationship info", http.StatusInternalServerError)
	}
}
