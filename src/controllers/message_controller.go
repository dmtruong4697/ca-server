package controllers

import (
	"ca-server/src/database"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
)

type GroupMessage struct {
	Sender  GetUserInfoResponce   `json:"sender"`
	Message models.Message        `json:"message"`
	Media   []models.MediaMessage `json:"media"`
}

type GroupMessages struct {
	Messages []GroupMessage
}

// get group's message
func GetGroupMessage(w http.ResponseWriter, r *http.Request) {
	current_user_id := r.Context().Value("id").(uint)

	var group_id uint
	err := json.NewDecoder(r.Body).Decode(&group_id)
	if err != nil {
		http.Error(w, "Failed to decode group id", http.StatusBadRequest)
	}

	// find all message
	var messages []models.Message
	if err := database.DB.Where("group_id = ?", group_id).Find(&messages).Error; err != nil {
		http.Error(w, "Message not found", http.StatusBadRequest)
		return
	}

	var group_messages GroupMessages
	for i := range messages {
		// find all message's media
		var medias []models.MediaMessage
		if err := database.DB.Where("message_id = ?", messages[i].ID).Find(&medias).Error; err != nil {
			http.Error(w, "Message media not found", http.StatusBadRequest)
			return
		}

		group_messages.Messages[i].Message = messages[i]
		group_messages.Messages[i].Sender = GetUserInfoFunction(current_user_id, medias[i].SenderID)
		group_messages.Messages[i].Media = medias
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(group_messages); err != nil {
		http.Error(w, "Failed to encode group info", http.StatusInternalServerError)
	}
}
