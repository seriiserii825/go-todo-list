package docker

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

func StartInDocker() {
	fmt.Println("PID", os.Getpid())

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	job(ctx)

	fmt.Println("Application stopped correctly")
}
