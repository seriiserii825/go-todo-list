package main

import (
	"context"
	"fmt"
	"time"
	"todo/db"
)

func main() {
	ctx := context.Background()
	conn, err := db.DbConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := db.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("Table created successfully")

	if err := db.InsertRow(
		ctx,
		conn,
		"Good Task 3",
		"This is a sample task description",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}
	fmt.Println("Row inserted successfully")
}
