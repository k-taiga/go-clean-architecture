package transaction

import (
	"context"
	"database/sql"
)

type Transaction interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error)
}

type tx struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &tx{db}
}

// DoInTx はトランザクション内で実行する処理を受け取りその処理の結果を返す
func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error) {
	// dbのトランザクションを開始する
	tx, err := t.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// contextにトランザクションをセットする
	ctx = context.WithValue(ctx, "tx", tx)

	// fにcontextを渡して実行する
	v, err := f(ctx)
	if err != nil {
		// fの処理が失敗したらロールバックする
		tx.Rollback()
		return nil, err
	}

	// fの処理が成功したらコミットする
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return v, nil
}
