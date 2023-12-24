package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"todoapp/model"
	"todoapp/usecase"

	"github.com/labstack/echo/v4"
)

type TaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
}

// interfaceを満たすtaskControllerを返す
func NewTaskController(u usecase.TaskUseCase) TaskController {
	// キーと値が同じ場合は省略できる{u:u}=>{u}
	return &taskController{u}
}

type taskController struct {
	u usecase.TaskUseCase
}

func (t *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("parse error: %v", err.Error())
		return c.JSON(http.StatusBadRequest, msg.Error())
	}
	task, err := t.u.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// c.JSONを返しているがechoでerror型の値を返している
	return c.JSON(http.StatusOK, task)
}

func (t *taskController) Create(c echo.Context) error {
	var task model.Task

	// c.BindはJSONを構造体にバインドする
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdID, err := t.u.CreateTask(task.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, createdID)
}
