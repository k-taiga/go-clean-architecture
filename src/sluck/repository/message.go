package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type MessageRepository interface {
	Delete(ctx context.Context, userID string) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) Delete(ctx context.Context, userID string) error {
	// 同じDBを使うことでトランザクション内で実行したSQLがatomicになる
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	_, err := db.Exec("DELETE FROM messages WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
