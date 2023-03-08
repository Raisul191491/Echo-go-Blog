package controllers

import (
	"fmt"
	domain "go-blog/pkg/domains"
	"go-blog/pkg/models"
	"go-blog/pkg/services"
	"go-blog/pkg/types"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var UserService domain.IUserService

func SetUserService(userService domain.IUserService) {
	UserService = userService
}

func Registration(e echo.Context) error {

	if os.Getenv("Auth") != "" && os.Getenv("ID") != "" && os.Getenv("Email") != "" {
		return e.JSON(http.StatusBadRequest, "First Log out of existing account")
	}

	user := &types.ControlUser{}
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

	tempUser := UserService.GetUser(newUser.Email)
	if len(tempUser) > 0 {
		return e.JSON(http.StatusBadRequest, "Account with this email already exists!")
	}

	err := UserService.RegisterUser(newUser)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "User not created")
	}
	return e.JSON(http.StatusCreated, "User created")
}

func Login(e echo.Context) error {
	if os.Getenv("Auth") != "" && os.Getenv("ID") != "" && os.Getenv("Email") != "" {
		return e.JSON(http.StatusBadRequest, "Already logged in to another account")
	}

	loginUser := &types.LoginType{}
	if err := e.Bind(loginUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	user := UserService.GetUser(loginUser.Email)
	if len(user) != 1 {
		return e.JSON(http.StatusBadRequest, "Account does not exist!")
	}

	if err := UserService.CheckPassword(loginUser.Password, user[0].Password); err != nil {
		return e.JSON(http.StatusUnauthorized, err.Error())
	}

	token, err := services.GenerateToken(user[0].Email, user[0].Username)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Could not generate token")
	}
	os.Setenv("Auth", token)
	os.Setenv("Email", loginUser.Email)
	os.Setenv("ID", fmt.Sprint(user[0].ID))

	return e.JSON(http.StatusOK, "Successfully logged in...")
}

func Logout(e echo.Context) error {
	os.Unsetenv("Email")
	os.Unsetenv("ID")
	os.Unsetenv("Auth")
	return e.JSON(http.StatusOK, "Successfully logged out")
}

func GetProfiles(e echo.Context) error {
	var tempUsers []models.User
	email := e.QueryParam("email")
	tempUsers = UserService.GetUser(email)
	users := UserService.RemoveSensitiveData(tempUsers)

	if len(users) == 0 {
		return e.JSON(http.StatusOK, "No user found")
	}

	if len(users) == 1 {
		blogList := services.GetBlogs(int(users[0].ID), 0)
		users[0].Blogs = blogList
	}

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

	id, err := GetIntEnv("ID")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Parsing error")
	}

	if err := Logout(e); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := services.DeleteBlog(0, id); err != nil {
		return e.JSON(http.StatusBadRequest, "Could not delete associate blogs")
	}

	if err := UserService.DeleteProfile(deleteProfile.Email); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "Account deleted...")

}

func UpdateProfile(e echo.Context) error {
	updateProfile := &types.ControlUser{}
	if err := e.Bind(updateProfile); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	currentProfile := UserService.GetUser(os.Getenv("Email"))[0]
	currentEmail := currentProfile.Email

	newProfile, validateProfile := ChangeProfileParams(*updateProfile, currentProfile)

	if err := validateProfile.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if currentEmail != newProfile.Email && len(UserService.GetUser(newProfile.Email)) > 0 {
		return e.JSON(http.StatusBadRequest, "Account with this email already exists!")
	}

	if err := UserService.UpdateProfile(&newProfile); err != nil {
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

func ChangeProfileParams(updateProfile types.ControlUser, currentProfile models.User) (models.User, types.ControlUser) {
	if updateProfile.Username != "" {
		currentProfile.Username = updateProfile.Username
	}
	if updateProfile.Email != "" {
		currentProfile.Email = updateProfile.Email
	}
	if updateProfile.Password != "" {
		currentProfile.Password = updateProfile.Password
	}

	tempProfile := types.ControlUser{
		Username: currentProfile.Username,
		Email:    currentProfile.Email,
		Password: currentProfile.Password,
	}

	return currentProfile, tempProfile
}
