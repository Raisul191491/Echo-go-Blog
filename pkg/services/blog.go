package services

import (
	"go-blog/pkg/models"
	"go-blog/pkg/types"
)

var BlogInterface models.IBlog

func SetBlogInterface(blog models.IBlog) {
	BlogInterface = blog
}

func CreateBlog(post *models.Blog) error {
	if createErr := BlogInterface.CreateBlog(post); createErr != nil {
		return createErr
	}
	return nil
}

func GetBlogs(userId, postId int) []types.CustomBlogResponse {
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

func DeleteBlog(postId, userId int) error {
	if deleteErr := BlogInterface.DeleteBlog(postId, userId); deleteErr != nil {
		return deleteErr
	}
	return nil
}
