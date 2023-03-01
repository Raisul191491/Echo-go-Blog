package controllers

import (
	"fmt"
	"go-blog/pkg/models"
	"go-blog/pkg/services"
	"go-blog/pkg/types"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Registration(e echo.Context) error {
	user := &models.User{}
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	if err := user.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	tempUser := services.Get(newUser.Email)
	if len(tempUser) > 0 {
		return e.JSON(http.StatusBadRequest, "Account with this email already exists!")
	}

	err := services.RegisterUser(newUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "User not created")
	}
	return e.JSON(http.StatusCreated, "User created")
}

func Login(e echo.Context) error {
	loginUser := &types.LoginType{}
	if err := e.Bind(loginUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	user := services.Get(loginUser.Email)
	if len(user) != 1 {
		return e.JSON(http.StatusBadRequest, "Account does not exist!")
	}

	if err := services.CheckPassword(loginUser.Password, user[0].Password); err != nil {
		return e.JSON(http.StatusUnauthorized, err.Error())
	}

	token, err := services.GenerateToken(user[0].Email, user[0].Username)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "err.Error()")
	}
	os.Setenv("Auth", token)
	os.Setenv("Email", loginUser.Email)
	os.Setenv("ID", fmt.Sprint(user[0].ID))

	return e.JSON(http.StatusOK, "Successfully logged in...")
}

func Logout(e echo.Context) error {
	if err := os.Setenv("Auth", ""); err != nil {
		return err
	}
	return e.JSON(http.StatusOK, "Successfully logged out")
}

func GetProfiles(e echo.Context) error {
	var tempUsers []models.User
	email := e.QueryParam("email")
	tempUsers = services.Get(email)
	users := services.RemoveSensitiveData(tempUsers)

	return e.JSON(http.StatusOK, users)
}

func DeleteProfile(e echo.Context) error {
	deleteProfile := &types.Deletetype{}
	if err := e.Bind(deleteProfile); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	if deleteProfile.Email != os.Getenv("Email") {
		return e.JSON(http.StatusBadRequest, "Not authorized to delete this account")
	}

	if err := Logout(e); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := services.DeleteProfile(deleteProfile.Email); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "Account deleted...")

}

func UpdateProfile(e echo.Context) error {
	updateProfile := &types.RegistrationType{}
	if err := e.Bind(updateProfile); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	currentProfile := services.Get(os.Getenv("Email"))[0]
	currentEmail := currentProfile.Email

	newProfile := ChangeProfileParams(*updateProfile, currentProfile)
	if err := newProfile.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if currentEmail != newProfile.Email && len(services.Get(newProfile.Email)) > 0 {
		return e.JSON(http.StatusBadRequest, "Account with this email already exists!")
	}

	if err := services.UpdateProfile(&newProfile); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	token, err := services.GenerateToken(newProfile.Email, newProfile.Username)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "err.Error()")
	}
	os.Setenv("Auth", token)
	os.Setenv("Email", newProfile.Email)

	return e.JSON(http.StatusCreated, "Successfull updated profile")
}

func ChangeProfileParams(updateProfile types.RegistrationType, currentProfile models.User) models.User {
	if updateProfile.Username != "" {
		currentProfile.Username = updateProfile.Username
	}
	if updateProfile.Email != "" {
		currentProfile.Email = updateProfile.Email
	}
	if updateProfile.Password != "" {
		currentProfile.Password = updateProfile.Password
	}

	return currentProfile
}
