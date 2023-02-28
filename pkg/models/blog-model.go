package models

import (
	"time"

	"github.com/labstack/echo"
)

type Blog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"userID"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
type INews interface {
	CreateBlog(e echo.Context) error
	UpdateBlog(e echo.Context) error
	GetAnyBlog(e echo.Context) error
	DeleteBlog(e echo.Context) error
}
