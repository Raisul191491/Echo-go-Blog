package domain

import "go-blog/pkg/models"

type IBlogRepo interface {
	CreateBlog(post *models.Blog) error
	GetAnyBlog(userId, postId int) []models.Blog
	DeleteBlog(postId, userId int) error
	// UpdateBlog(e echo.Context) error
}
