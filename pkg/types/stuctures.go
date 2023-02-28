package types

import "github.com/golang-jwt/jwt"

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
