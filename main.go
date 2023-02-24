package main

import (
	"go-blog/pkg/containers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Init(e)
}
