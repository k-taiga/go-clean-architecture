package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (controller *TaskController) Get(c echo.Context) error {
	// tasks,err := usecase.GetTasks()
	return c.JSON(http.StatusOK, nil)
}

func (controller *TaskController) Create(c echo.Context) error {
	var task Task

	// c.BindはJSONを構造体にバインドする
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	// created, err := usecase.CreateTask(&task)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, nil)
	// }

	return c.JSON(http.StatusOK, nil)
}
