package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1,$2, $3, $4);
	`

	if task.CreatedAt.IsZero() {
		task.CreatedAt = time.Now()
	}

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt)
	return err
}
