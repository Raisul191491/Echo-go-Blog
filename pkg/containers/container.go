package containers

import (
	"go-blog/pkg/routes"
	"go-blog/pkg/utils"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	utils.Connect()

	// Ibookcrud := repositories.BookDbInstance(db)
	// Iauthorcrud := repositories.AuthorDbInstance(db)
	// services.BookInterfaceInstance(Ibookcrud)
	// services.AuthorInterfaceInstance(Iauthorcrud)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// routes.RegisteredBookStoreRoutes(r)
	routes.UserBlogRoutes(e)
	e.Start(":9010")
}
