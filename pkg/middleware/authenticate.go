package middleware

import (
	"go-blog/pkg/controllers"
	"go-blog/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cacheMap, _ := controllers.CheckCache()
		token := cacheMap["Auth"]
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Login required...")
		}

		if verifyErr := services.VerifyToken(token); verifyErr != nil {
			return c.JSON(http.StatusUnauthorized, verifyErr.Error())
		}

		return next(c)
	}
}
