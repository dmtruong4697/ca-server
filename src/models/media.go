package models

import (
	"time"
)

type MediaMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	SenderID  uint      `json:"sender_id"`
	MessageID uint      `json:"message_id"`
	Type      string    `json:"type"`
}
