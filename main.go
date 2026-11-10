package main

import (
	"context"
	"fmt"
	"todo/db"

	"github.com/joho/godotenv"
)

func main() {
	// val := os.Getenv("phone_number")
	// if val != "" {
	// 	fmt.Println("val: ", val)
	// } else {
	// 	fmt.Println("val is empty")
	// }

	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file, using system env")
	}

	ctx := context.Background()
	conn, err := db.DbConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	// tasks, err := db.SelectPagination(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, task := range tasks {
	// 	if task.ID == 20 {
	// 		task.Title = "Updated Task Title 20"
	// 		task.Description = "Updated Task Description 20"
	// 		task.Completed = true
	//
	// 		if err := db.UpdateTask(ctx, conn, task); err != nil {
	// 			panic(err)
	// 		}
	// 		pp.Println("Task updated successfully")
	// 		pp.Println(tasks)
	// 		return
	// 	}
	// }
	//
	// if err := db.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Table created successfully")
	//
	// if err := db.UpdateRow(ctx, conn, db.TaskModel{ID: 20, Completed: true}); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Row updated successfully")
	//
	// newTask := db.TaskModel{
	// 	Title:       "Good Task 5",
	// 	Description: "This is a sample task description",
	// 	Completed:   false,
	// 	CreatedAt:   time.Now(),
	// }
	//
	// if err := db.InsertRow(
	// 	ctx,
	// 	conn,
	// 	newTask,
	// ); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Row inserted successfully")
}
