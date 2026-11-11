package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func job(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Job stopped")
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("Job running... %d\n", i)
			i++
		}
	}
}

func main() {
	fmt.Println("PID", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	job(ctx)

	fmt.Println("Application stopped correctly")

	// fmt.Println("Starting the HTTP server...")
	// err := http_server.StartHTTPServer()
	// if err != nil {
	// 	panic(err)
	// }

	// val := os.Getenv("phone_number")
	// if val != "" {
	// 	fmt.Println("val: ", val)
	// } else {
	// 	fmt.Println("val is empty")
	// }

	// ctx := context.Background()
	// conn, err := db.DbConnection(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close(ctx)

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
