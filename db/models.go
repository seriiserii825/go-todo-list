package db

import "time"

type TaskModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
