package services

import "go-blog/pkg/models"

var UserInterface models.IUser

func SetUserInterface(user models.IUser) {
	UserInterface = user
}
