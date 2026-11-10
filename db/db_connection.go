package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DbConnection(ctx context.Context) (*pgx.Conn, error) {
	conn_str := os.Getenv("CONN_STR")
	fmt.Println("conn_str: ", conn_str)
	connect_pg_string := conn_str
	return pgx.Connect(ctx, connect_pg_string)
}
