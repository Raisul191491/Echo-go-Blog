package types

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type NewBlogBody struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type CustomBlogResponse struct {
	ID        uint      `json:"id,omitempty"`
	Subject   string    `json:"subject,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserID    uint      `json:"userID,omitempty"`
	Username  string    `json:"username,omitempty"`
}

type ControlBlog struct {
	ID        uint        `json:"id"`
	Subject   string      `json:"subject"`
	Body      string      `json:"body"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	UserID    uint        `json:"userID"`
	User      ControlUser `json:"user"`
}

func (b ControlBlog) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Subject,
			validation.Required.Error("Please input subject of your post"),
			validation.Length(5, 60)),
		validation.Field(&b.Body,
			validation.Required.Error("Description needed!"),
			validation.Length(6, 300)),
	)
}
