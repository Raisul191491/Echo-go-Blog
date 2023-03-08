package domain

import (
	"go-blog/pkg/models"
	"go-blog/pkg/types"
)

type IBlogRepo interface {
	CreateBlog(post *models.Blog) error
	GetAnyBlog(userId, postId int) []models.Blog
	DeleteBlog(postId, userId int) error
	// UpdateBlog(e echo.Context) error
}

type IBlogService interface {
	CreateBlog(post *models.Blog) error
	GetBlogs(userId, postId int) []types.CustomBlogResponse
	DeleteBlog(postId, userId int) error
}
