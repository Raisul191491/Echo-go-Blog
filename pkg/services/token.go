package services

import (
	domain "go-blog/pkg/domains"
)

var TokenAuth domain.IToken

func SetTokenAuth(tAuth domain.IToken) {
	TokenAuth = tAuth
}

func GenerateToken(email, username string) (string, error) {
	tokenString, err := TokenAuth.GenerateJWT(email, username)

	return tokenString, err
}

func VerifyToken(token string) error {
	if err := TokenAuth.ValidateToken(token); err != nil {
		return err
	}
	return nil
}
