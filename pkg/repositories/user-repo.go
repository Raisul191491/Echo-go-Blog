package repositories

import (
	"go-blog/pkg/models"

	"gorm.io/gorm"
)

var db *gorm.DB

type dbUser struct {
	DB *gorm.DB
}

func UserDBInstance(d *gorm.DB) models.IUser {
	db = d
	return &dbUser{
		DB: db,
	}
}

func (repo *dbUser) Register(user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbUser) GetUsers(email string) []models.User {
	var users []models.User
	if email == "" {
		db.Find(&users)
	} else {
		db.Where("email = ?", email).Find(&users)
	}
	return users
}

func (repo *dbUser) DeleteProfile(email string) error {
	var deletedUser models.User
	if err := db.Where("email = ?", email).Delete(&deletedUser).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbUser) UpdateProfile(user *models.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
