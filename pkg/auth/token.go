package auth

import (
	"errors"
	domain "go-blog/pkg/domains"
	"go-blog/pkg/types"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenAuth struct {
	tRepo domain.IUserRepo
}

func TokenAuthInstance(uRepo domain.IUserRepo) domain.IToken {
	return &TokenAuth{
		tRepo: uRepo,
	}
}

func (t *TokenAuth) GenerateJWT(email, username string) (string, error) {
	claim := &types.Claim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("raisul-islam"))
	return tokenString, err
}

func (t *TokenAuth) ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&types.Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("raisul-islam"), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*types.Claim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return nil
}
