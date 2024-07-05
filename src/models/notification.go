package models

import (
	"time"
)

type Notification struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    bool      `json:"is_read"`
	Content   string    `json:"content"`
}
