package models

import (
	"time"
)

type Group struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	CreatorID         uint      `json:"creator_id"`
	Name              string    `json:"name"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	LastMessageID     uint      `json:"last_message_id"`
	GroupImageURL     string    `json:"group_image_url"`
	Type              string    `json:"type"`
	InviteCode        string    `json:"invite_code"`
	IsAllowInviteCode bool      `json:"is_allow_invite_code"`
}
