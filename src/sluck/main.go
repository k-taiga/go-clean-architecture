package main

import (
	"net/http"
	"sluck/controller"
	"sluck/infra"
	"sluck/repository"
	"sluck/transaction"
	"sluck/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	// validator.Struct(i) は i anyの構造体に対してバリデーションを行う
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	// カスタムバリデータを登録する
	e.Validator = &CustomValidator{validator: validator.New()}

	db := infra.ConnectDB()

	ur := repository.NewUserRepository(db)
	mr := repository.NewMessageRepository(db)
	uu := usecase.NewUserUsecase(ur, mr, transaction.NewTransaction(db))
	uc := controller.NewUserController(uu)

	e.POST("/users", uc.Create)
	// getとdeleteはidをパラメータに取る(ctx.Param("id")で取得できる)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users", uc.Update)
	e.DELETE("/users/:id", uc.Delete)
	e.Start(":8080")
}
