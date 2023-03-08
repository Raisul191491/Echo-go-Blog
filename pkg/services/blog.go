package services

import (
	domain "go-blog/pkg/domains"
	"go-blog/pkg/models"
	"go-blog/pkg/types"
)

var BlogInterface domain.IBlogRepo

type BlogService struct {
	bRepo domain.IBlogRepo
}

func BlogServiceInstance(blogRepo domain.IBlogRepo) domain.IBlogService {
	return &BlogService{
		bRepo: blogRepo,
	}
}

func SetBlogInterface(blog domain.IBlogRepo) {
	BlogInterface = blog
}

func (b *BlogService) CreateBlog(post *models.Blog) error {
	if createErr := BlogInterface.CreateBlog(post); createErr != nil {
		return createErr
	}
	return nil
}

func (b *BlogService) GetBlogs(userId, postId int) []types.CustomBlogResponse {
	var finalList []types.CustomBlogResponse
	blogList := BlogInterface.GetAnyBlog(userId, postId)
	for _, val := range blogList {
		finalList = append(finalList, types.CustomBlogResponse{
			ID:        val.ID,
			Subject:   val.Subject,
			Body:      val.Body,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			UserID:    val.UserID,
			Username:  val.User.Username,
		})
	}
	return finalList
}

func (b *BlogService) DeleteBlog(postId, userId int) error {
	if deleteErr := BlogInterface.DeleteBlog(postId, userId); deleteErr != nil {
		return deleteErr
	}
	return nil
}
