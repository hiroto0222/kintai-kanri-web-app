package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store はデータベースに対してクエリを実行するための基本的な関数Querierを持つ構造体
type Store interface {
	Querier
	ClockOutTx(ctx context.Context, arg ClockOutTxParams) (ClockOutTxResult, error)
}

// SQLStore はsql.DBを持つ構造体
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
