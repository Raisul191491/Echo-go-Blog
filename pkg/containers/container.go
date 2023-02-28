package containers

import (
	"go-blog/pkg/repositories"
	"go-blog/pkg/routes"
	"go-blog/pkg/services"
	"go-blog/pkg/utils"
	"log"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	utils.Connect()
	db := utils.GetDB()
	userInterface := repositories.UserDBInstance(db)
	services.SetUserInterface(userInterface)
	routes.UserBlogRoutes(e)
	log.Fatal(e.Start(":9020"))
}
