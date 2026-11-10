package main

import (
	"context"
	"fmt"
	"todo/db"
)

func main() {
	ctx := context.Background()
	conn, error := db.DbConnection(ctx)
	if error != nil {
		panic(error)
	}
	if db.CreateTable(ctx, conn); error != nil {
		panic(error)
	}
	fmt.Println("Table created successfully")
}
