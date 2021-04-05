package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				fmt.Println(counter.Count()) // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

// 使用读锁保护
func (c *Counter) Count() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.count
}
