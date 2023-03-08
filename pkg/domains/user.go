package domain

import (
	"go-blog/pkg/models"
	"go-blog/pkg/types"
)

type IUserRepo interface {
	RegisterUser(user *models.User) error
	GetUsers(email string) []models.User
	UpdateProfile(user *models.User) error
	DeleteProfile(email string) error
}

type IUserService interface {
	RegisterUser(user *models.User) error
	GetUser(email string) []models.User
	DeleteProfile(email string) error
	UpdateProfile(user *models.User) error
	CheckPassword(loginPass, hashedPass string) error
	RemoveSensitiveData(users []models.User) []types.CustomProfileResponse
}
