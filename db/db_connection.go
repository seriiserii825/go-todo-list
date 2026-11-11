package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func DbConnection(ctx context.Context) (*pgx.Conn, error) {
	conn_str := os.Getenv("CONN_STR")
	connect_pg_string := conn_str
	return pgx.Connect(ctx, connect_pg_string)
}
