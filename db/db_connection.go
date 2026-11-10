package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DbConnection() {
	ctx := context.Background()
	connect_pg_string := "postgres://postgres:postgres@localhost:5434/todo"
	conn, err := pgx.Connect(ctx, connect_pg_string)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database successfully!")
}
