package models

import "time"

type GroupMember struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	GroupID     uint      `json:"group_id"`
	JoinAt      time.Time `json:"join_at"`
	Status      string    `json:"status"`
	IngroupName string    `json:"ingroup_name"`
}
