package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Blog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"userID"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
type IBlog interface {
	CreateBlog(post *Blog) error
	GetAnyBlog(userId, postId int) []Blog
	DeleteBlog(postId, userId int) error
	// UpdateBlog(e echo.Context) error
}

func (b Blog) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Subject,
			validation.Required.Error("Please input subject of your post"),
			validation.Length(5, 60)),
		validation.Field(&b.Body,
			validation.Required.Error("Description needed!"),
			validation.Length(6, 300)),
	)
}
