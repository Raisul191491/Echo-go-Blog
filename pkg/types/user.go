package types

import (
	"time"
)

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Deletetype struct {
	Email string `json:"email"`
}

type RegistrationType struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomProfileResponse struct {
	ID        uint                 `json:"id,omitempty"`
	Username  string               `json:"username,omitempty"`
	Email     string               `json:"email,omitempty"`
	UpdatedAt time.Time            `json:"updated_at,omitempty"`
	Blogs     []CustomBlogResponse `json:"blogs,omitempty"`
}

type ControlUser struct {
	ID               uint      `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
