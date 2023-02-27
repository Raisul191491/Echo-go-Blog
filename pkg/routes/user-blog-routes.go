package routes

import (
	"go-blog/pkg/controllers"
	"go-blog/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserBlogRoutes(e *echo.Echo) {
	user := e.Group("/user", middleware.Authenticate)

	/* User Routes */
	user.POST("/registration", controllers.Registration)
	user.POST("/login", controllers.Login)
	user.GET("/:id", controllers.GetProfiles)
	user.PUT("/:id", controllers.UpdateProfile)
	user.DELETE("/:id", controllers.DeleteProfile)

	/* News Routes */
	user.POST("blog", controllers.CreateBlog)
	user.GET("blog", controllers.GetAnyBlog)
	user.PUT("/blog/:id", controllers.UpdateBlog)
	user.PUT("blog/:id", controllers.DeleteBlog)
}
