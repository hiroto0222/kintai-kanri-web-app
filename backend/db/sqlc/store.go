package db

import (
	"database/sql"
)

// Store はデータベースに対してクエリを実行するための基本的な関数Querierを持つ構造体
type Store interface {
	Querier
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
