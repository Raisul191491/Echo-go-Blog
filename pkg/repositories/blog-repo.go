package repositories

import (
	"go-blog/pkg/models"

	"gorm.io/gorm"
)

type dbBlog struct {
	DB *gorm.DB
}

func BlogDBInstance(d *gorm.DB) models.IBlog {
	db = d
	return &dbBlog{
		DB: db,
	}
}

func (repo *dbBlog) CreateBlog(post *models.Blog) error {
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (repo *dbBlog) GetAnyBlog(userId, postId int) []models.Blog {
	var blogs []models.Blog
	if postId != 0 {
		db.Joins("User").Where("`blogs`.`id` = ?", postId).Find(&blogs)
		return blogs
	}
	if userId != 0 {
		db.Joins("User").Where("user_id = ?", userId).Find(&blogs)
		return blogs
	}
	db.Joins("User").Find(&blogs)
	return blogs
}

func (repo *dbBlog) DeleteBlog(postId int) error {
	var deletedBlog models.Blog
	if err := db.Where("id = ?", postId).Delete(deletedBlog).Error; err != nil {
		return err
	}
	return nil
}
