package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectPagination(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sql_row := `SELECT * FROM tasks ORDER BY id LIMIT 10 OFFSET 0;`

	rows, err := conn.Query(ctx, sql_row)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	tasks := []TaskModel{}

	for rows.Next() {
		var task TaskModel

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}
