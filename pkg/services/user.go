package services

import (
	"errors"
	domain "go-blog/pkg/domains"
	"go-blog/pkg/models"
	"go-blog/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

var UserInterface domain.IUserRepo

type UserService struct {
	uRepo domain.IUserRepo
}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &UserService{
		uRepo: userRepo,
	}
}

func SetUserInterface(user domain.IUserRepo) {
	UserInterface = user
}

func (u *UserService) RegisterUser(user *models.User) error {

	tempPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = tempPassword

	regErr := UserInterface.RegisterUser(user)
	return regErr
}

func (u *UserService) GetUser(email string) []models.User {
	userlist := UserInterface.GetUsers(email)
	return userlist
}

func (u *UserService) DeleteProfile(email string) error {
	deleteErr := UserInterface.DeleteProfile(email)
	return deleteErr
}

func (u *UserService) UpdateProfile(user *models.User) error {
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

func (u *UserService) CheckPassword(loginPass, hashedPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(loginPass)); err != nil {
		return errors.New("wrong credential(password)")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserService) RemoveSensitiveData(users []models.User) []types.CustomProfileResponse {
	var finalUsers []types.CustomProfileResponse
	for _, val := range users {
		finalUsers = append(finalUsers, types.CustomProfileResponse{
			ID:        val.ID,
			Username:  val.Username,
			Email:     val.Email,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return finalUsers
}
