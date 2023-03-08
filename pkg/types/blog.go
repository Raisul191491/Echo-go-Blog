package types

import "time"

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
