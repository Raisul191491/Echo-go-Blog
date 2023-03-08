package domain

import (
	"go-blog/pkg/models"
)

type IUserRepo interface {
	Register(user *models.User) error
	GetUsers(email string) []models.User
	UpdateProfile(user *models.User) error
	DeleteProfile(email string) error
}
