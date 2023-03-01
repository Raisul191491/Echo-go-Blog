package services

import (
	"errors"
	"go-blog/pkg/models"
	"go-blog/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	tempPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = tempPassword

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
	tempPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = tempPassword

	if updateErr := UserInterface.UpdateProfile(user); updateErr != nil {
		return updateErr
	}
	return nil
}

func CheckPassword(loginPass, hashedPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(loginPass)); err != nil {
		return errors.New("wrong credential(password)")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RemoveSensitiveData(users []models.User) []types.CustomResponse {
	var finalUsers []types.CustomResponse
	for _, val := range users {
		finalUsers = append(finalUsers, types.CustomResponse{
			ID:        val.ID,
			Username:  val.Username,
			Email:     val.Email,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return finalUsers
}
