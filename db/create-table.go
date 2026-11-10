package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
