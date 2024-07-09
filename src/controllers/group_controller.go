package controllers

import (
	"ca-server/src/database"
	"ca-server/src/models"
	"encoding/json"
	"net/http"
	"time"
)

type LastMessage struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	Content   string              `json:"content"`
	Sender    GetUserInfoResponce `json:"sender"`
}

type Member struct {
	UserInfo    GetUserInfoResponce `json:"user_info"`
	JoinAt      time.Time           `json:"join_at"`
	Status      string              `json:"status"`
	IngroupName string              `json:"ingroup_name"`
}

type GetGroupInfoResponce struct {
	Group       models.Group `json:"group"`
	Members     []Member     `json:"members"`
	LastMessage LastMessage  `json:"last_message"`
}

// get group detail
func GetGroupInfo(w http.ResponseWriter, r *http.Request) {
	current_user_id := r.Context().Value("id").(uint)

	var dbUser models.User
	if err := database.DB.Where("id = ?", current_user_id).First(&dbUser).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// get group id from request body
	var group_id uint
	err := json.NewDecoder(r.Body).Decode(&group_id)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusBadRequest)
		return
	}

	// check user in group or not
	var group_member models.GroupMember
	if err := database.DB.Where("(user_id = ? AND group_id = ?)", current_user_id, group_id).First(&group_member).Error; err != nil {
		http.Error(w, "User is not Group member", http.StatusBadRequest)
		return
	}

	// get group from database
	var dbGroup models.Group
	if err := database.DB.Where("id = ?", group_id).First(&dbGroup).Error; err != nil {
		http.Error(w, "Group not found", http.StatusBadRequest)
		return
	}

	// get all member id
	var group_members []models.GroupMember
	if err := database.DB.Where("group_id = ?", group_id).Find(&group_members).Error; err != nil {
		http.Error(w, "Group member not found", http.StatusBadRequest)
		return
	}

	// get group members detail
	var group_members_detail []Member
	for i := range group_members {
		group_members_detail[i].UserInfo = GetUserInfoFunction(current_user_id, group_members[i].UserID)
		group_members_detail[i].JoinAt = group_member.JoinAt
		group_members_detail[i].Status = group_member.Status
		group_members_detail[i].IngroupName = group_member.IngroupName
	}

	// get last message detail
	var dbMessage models.Message
	if err := database.DB.Where("id = ?", dbGroup.LastMessageID).First(&dbMessage).Error; err != nil {

	}
	var last_message LastMessage
	last_message.ID = dbMessage.ID
	last_message.CreatedAt = dbMessage.CreatedAt
	last_message.Content = dbMessage.Content
	last_message.Sender = GetUserInfoFunction(current_user_id, dbMessage.SenderID)

	// init responce
	var responce GetGroupInfoResponce
	responce.Group = dbGroup
	responce.Members = group_members_detail
	responce.LastMessage = last_message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responce); err != nil {
		http.Error(w, "Failed to encode responce info", http.StatusInternalServerError)
	}
}
