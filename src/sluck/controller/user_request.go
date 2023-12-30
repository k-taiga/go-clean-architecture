package controller

import (
	"sluck/model"
	"time"
)

type UserRequest struct {
	// validate:"required" は必須項目であることを示す(validatorで定義されているものが使える)
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// requestをmodelにバインドする
// 構造体をポインタ型で渡したほうがメモリ効率が良い
// ポインタ型はアドレスを格納するための変数 ポインタ型は * で表す model.Userの構造体のアドレスを格納するための変数
func toModel(req UserRequest) *model.User {
	now := time.Now()
	// &model.User は User のアドレスを返す ex. 0xc0000b4000
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
