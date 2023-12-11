package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World from echo")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.String(200, "Hello Users")
	})

	e.Start(":8080")
}
