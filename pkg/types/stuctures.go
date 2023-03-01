package types

import (
	"time"

	"github.com/golang-jwt/jwt"
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

type Claim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type CustomResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewBlogBody struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type CustomBlogResponse struct {
	ID        uint      `json:"id"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"userID"`
	Username  string    `json:"username"`
}
