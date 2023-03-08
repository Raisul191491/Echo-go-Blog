package types

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Deletetype struct {
	Email string `json:"email"`
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

func (u ControlUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username,
			validation.Required.Error("Enter user name"),
			validation.Length(5, 30)),
		validation.Field(&u.Email,
			validation.Required.Error("Email field cannot be empty"),
			is.Email),
		validation.Field(&u.Password,
			validation.Required.Error("password field cannot be empty"),
			validation.Length(8, 30)),
	)
}
