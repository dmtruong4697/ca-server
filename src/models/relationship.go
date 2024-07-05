package models

import (
	"time"
)

type Relationship struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
}
