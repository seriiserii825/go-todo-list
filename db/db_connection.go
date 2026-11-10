package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DbConnection(ctx context.Context) (*pgx.Conn, error) {
	connect_pg_string := "postgres://postgres:postgres@localhost:5434/todo"
	return pgx.Connect(ctx, connect_pg_string)
}
