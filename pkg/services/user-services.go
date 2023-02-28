package services

import (
	"errors"
	"go-blog/pkg/models"
)

func RegisterUser(user *models.User) error {
	regErr := UserInterface.Register(user)
	return regErr
}

func Get(email string) []models.User {
	userlist := UserInterface.GetUsers(email)
	return userlist
}

func DeleteProfile(email string) error {
	deleteErr := UserInterface.DeleteProfile(email)
	return deleteErr
}

func UpdateProfile(user *models.User) error {
	if updateErr := UserInterface.UpdateProfile(user); updateErr != nil {
		return updateErr
	}
	return nil
}

func CheckPassword(logPass, userPass string) error {
	if logPass == userPass {
		return nil
	}
	return errors.New("wrong credential(password)")
}
