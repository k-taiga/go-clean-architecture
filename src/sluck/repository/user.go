package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sluck/model"
	"strconv"
)

type UserRepository interface {
	Read(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Read(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id,name,email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	result, err := r.db.Exec("INSERT INTO users (name, age, email) VALUES (?, ?, ?)", &user.Name, &user.Age, &user.Email)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err

	}

	// id,10は10進数でidを文字列に変換する
	idStr := strconv.FormatInt(id, 10)
	return idStr, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	// userモデルから値を取り出して、UPDATE文を実行する
	result, err := r.db.Exec("UPDATE users SET name = ?, age = ?, email = ? WHERE id = ?", &user.Name, &user.Age, &user.Email, &user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", user.ID)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	// ctxからトランザクションを取り出す
	// transactionで使っているDBを使う
	// 同じDBを使うことでトランザクション内で実行したSQLがatomicになる
	db, ok := GetTx(ctx)

	if !ok {
		return fmt.Errorf("no transaction found")
	}

	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", id)
	}

	return nil
}
