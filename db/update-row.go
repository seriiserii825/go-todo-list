package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sql_row := `UPDATE tasks SET completed = $1 WHERE id = $2;`
	_, err := conn.Exec(ctx, sql_row, task.Completed, task.ID)
	return err
}

func UpdateTask(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sql_row := `UPDATE tasks SET title = $1, description = $2, completed = $3, updated_at = NOW() WHERE id = $4;`
	_, err := conn.Exec(ctx, sql_row, task.Title, task.Description, task.Completed, task.ID)
	return err
}
