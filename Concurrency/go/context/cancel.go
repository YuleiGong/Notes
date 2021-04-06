package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 父context
	ctx, cancel := context.WithCancel(context.Background())
	go watch1(ctx)

	// 子context
	valueCtx, _ := context.WithCancel(ctx)
	go watch2(valueCtx)

	time.Sleep(3 * time.Second)

	cancel()

	// 再等待5秒看输出，可以发现父context的子协程和子context的子协程都会被结束掉
	time.Sleep(5 * time.Second)
}

// 父context的协程
func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("parent goroutine exit")
			return
		default:
			fmt.Println("parent goroutine pending")
			time.Sleep(1 * time.Second)
		}
	}
}

// 子context的协程
func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child goroutine exit")
			return
		default:
			fmt.Println("child pending")
			time.Sleep(1 * time.Second)
		}
	}
}
