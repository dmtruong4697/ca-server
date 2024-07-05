package models

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SenderID   uint      `json:"sender_id"`
	GroupID    uint      `json:"group_id"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	IsEdited   bool      `json:"is_edited"`
	LastUpdate time.Time `json:"last_update"`
	Status     string    `json:"status"`
}
