package main

import (
	"database/sql"
	"log"
	"todoapp/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

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
