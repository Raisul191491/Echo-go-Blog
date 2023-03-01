package routes

import (
	"go-blog/pkg/controllers"
	"go-blog/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func UserBlogRoutes(e *echo.Echo) {
	user := e.Group("/user", middleware.Authenticate)
	app := e.Group("/app")

	/* App routes */
	app.POST("/registration", controllers.Registration)
	app.POST("/login", controllers.Login)

	/* User Routes */
	user.POST("/logout", controllers.Logout)
	user.GET("/profiles", controllers.GetProfiles)
	user.DELETE("/deleteprofile", controllers.DeleteProfile)
	user.PUT("/updateprofile", controllers.UpdateProfile)

	// /* News Routes */
	// user.POST("blog", controllers.CreateBlog)
	// user.GET("blog", controllers.GetAnyBlog)
	// user.PUT("/blog/:id", controllers.UpdateBlog)
	// user.PUT("blog/:id", controllers.DeleteBlog)
}
