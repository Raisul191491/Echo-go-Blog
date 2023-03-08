package containers

import (
	"go-blog/pkg/auth"
	"go-blog/pkg/connection"
	"go-blog/pkg/controllers"
	"go-blog/pkg/repositories"
	"go-blog/pkg/routes"
	"go-blog/pkg/services"
	"log"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	connection.Connect()
	db := connection.GetDB()

	userRepo := repositories.UserDBInstance(db)
	services.SetUserInterface(userRepo)
	userService := services.UserServiceInstance(userRepo)
	controllers.SetUserService(userService)

	blogRepo := repositories.BlogDBInstance(db)
	services.SetBlogInterface(blogRepo)
	blogService := services.BlogServiceInstance(blogRepo)
	controllers.SetBlogService(blogService)

	tokenAuth := auth.TokenAuthInstance(userRepo)
	services.SetTokenAuth(tokenAuth)

	routes.UserBlogRoutes(e)
	log.Fatal(e.Start(":9020"))
}
