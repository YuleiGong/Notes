package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))

	go watch(ctx, "watch1")
	go watch(ctx, "watch2")

	time.Sleep(8 * time.Second)

	fmt.Println("cancel")
	cancel()

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Deadline())
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
