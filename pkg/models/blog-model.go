package models

import (
	"time"
)

type News struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"userID"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
