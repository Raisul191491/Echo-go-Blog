package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

type IUser interface {
	Register(user *User) error
	GetUsers(email string) []User
	UpdateProfile(user *User) error
	DeleteProfile(email string) error
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username,
			validation.Required,
			validation.Length(5, 30)),
		validation.Field(&u.Email,
			validation.Required,
			is.Email),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 30)),
	)
}
