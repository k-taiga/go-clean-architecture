package main

import (
	"database/sql"
	"log"
	"todoapp/controller"
	"todoapp/repository"
	"todoapp/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3のドライバ
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

	// テーブル作成
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// log
	e.Use(middleware.Logger())
	// recover(panicを起こしたときに500エラーを返す)
	e.Use(middleware.Recover())

	// 依存関係を注入する
	repository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(repository)
	taskController := controller.NewTaskController(taskUseCase)

	e.GET("/tasks", taskController.Get)
	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}
