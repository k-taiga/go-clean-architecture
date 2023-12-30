package repository

import (
	"context"
	"database/sql"
)

var txKey = "tx"

func GetTx(ctx context.Context) (*sql.Tx, bool) {
	// contextから与えられたキーをもとにtransactionを取得する
	tx, ok := ctx.Value(txKey).(*sql.Tx)

	return tx, ok
}
