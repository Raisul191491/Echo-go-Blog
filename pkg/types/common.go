package types

import "github.com/golang-jwt/jwt"

type Claim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
