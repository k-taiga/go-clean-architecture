package main

import (
	"todoapp/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// log
	e.Use(middleware.Logger())
	// recover(panicを起こしたときに500エラーを返す)
	e.Use(middleware.Recover())

	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}
