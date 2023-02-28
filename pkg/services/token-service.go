package services

import (
	"go-blog/pkg/auth"
)

func GenerateToken(email, username string) (string, error) {
	tokenString, err := auth.GenerateJWT(email, username)

	return tokenString, err
}

func VerifyToken(token string) error {
	if err := auth.ValidateToken(token); err != nil {
		return err
	}
	return nil
}
