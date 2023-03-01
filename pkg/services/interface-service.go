package services

import "go-blog/pkg/models"

var UserInterface models.IUser
var BlogInterface models.IBlog

func SetUserInterface(user models.IUser) {
	UserInterface = user
}

func SetBlogInterface(blog models.IBlog) {
	BlogInterface = blog
}
