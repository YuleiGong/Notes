package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个子节点的context,3秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	go watch(ctx, "watch1")
	go watch(ctx, "watch2")

	time.Sleep(8 * time.Second)

	fmt.Println("cancel")
	cancel()

}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
