package models

import (
	"time"
)

type User struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	PhoneNumber   string    `json:"phone_number"`
	AvatarImage   string    `json:"avatar_image"`
	HashtagName   string    `json:"hashtag_name"`
	CreatedAt     time.Time `json:"created_at"`
	AccountStatus string    `json:"account_status"`
	ValidateCode  string    `json:"validate_code"`
	LastActive    time.Time `json:"last_active"`
	Gender        string    `json:"gender"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	DeviceToken   string    `json:"device_token"`
}
