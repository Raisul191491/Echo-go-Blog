package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Username         string    `gorm:"type:varchar(255);not null" json:"username"`
	Email            string    `gorm:"type:varchar(255);not null" json:"email"`
	Password         string    `gorm:"not null" json:"password"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// var testDB *gorm.DB

func SetTestDBInstance(tdb *gorm.DB) {
	_ = tdb
}
