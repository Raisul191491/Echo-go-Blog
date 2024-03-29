package controllers

import (
	"errors"
	domain "go-blog/pkg/domains"
	"go-blog/pkg/models"
	"go-blog/pkg/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var BlogService domain.IBlogService

func SetBlogService(blogService domain.IBlogService) {
	BlogService = blogService
}

func CreateBlog(e echo.Context) error {
	newBlog := &types.ControlBlog{}
	if err := e.Bind(newBlog); err != nil {
		return e.JSON(http.StatusBadRequest, "Bad inputs!")
	}

	id, err := GetIntEnv("ID")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "err.Error()")
	}

	blog := &models.Blog{
		Subject: newBlog.Subject,
		Body:    newBlog.Body,
		UserID:  uint(id),
	}

	if err := newBlog.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := BlogService.CreateBlog(blog); err != nil {
		return e.JSON(http.StatusInternalServerError, "Post not created")
	}

	return e.JSON(http.StatusCreated, "Post created successfully")
}

func GetAnyBlog(e echo.Context) error {
	tempBlog := e.QueryParam("blogId")
	tempUser := e.QueryParam("userId")

	blogId, blogErr := strconv.ParseInt(tempBlog, 0, 0)
	userId, userErr := strconv.ParseInt(tempUser, 0, 0)

	if tempBlog == "" && blogErr != nil {
		blogId = 0
	}
	if tempUser == "" && userErr != nil {
		userId = 0
	}

	blogs := BlogService.GetBlogs(int(blogId), int(userId))

	if len(blogs) == 0 {
		return e.JSON(http.StatusOK, "No post found")
	}

	return e.JSON(http.StatusOK, blogs)
}

// func UpdateBlog(e echo.Context) error
func DeleteBlog(e echo.Context) error {
	deleteId := e.Param("id")
	postId, postErr := strconv.ParseInt(deleteId, 0, 0)
	if postErr != nil {
		return e.JSON(http.StatusBadRequest, "Enter valid post ID")
	}

	checkBlog := BlogService.GetBlogs(0, int(postId))

	if len(checkBlog) != 1 {
		return e.JSON(http.StatusOK, "Post does not exist!")
	}

	userId, err := GetIntEnv("ID")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	if userId != int(checkBlog[0].UserID) {
		return e.JSON(http.StatusBadRequest, "Not authorized to delete this post")
	}

	if err := BlogService.DeleteBlog(int(postId), 0); err != nil {
		return e.JSON(http.StatusInternalServerError, "Could not delete post")
	}

	return e.JSON(http.StatusOK, "Successfully deleted post")
}

func GetIntEnv(key string) (int, error) {
	CacheMap, err := CheckCache()
	if err != nil {
		return 0, errors.New("caching error")
	}
	ret, err := strconv.Atoi(CacheMap["ID"])
	return ret, err
}
