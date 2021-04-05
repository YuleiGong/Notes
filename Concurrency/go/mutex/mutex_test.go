package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMetuxOne(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	t.Log(count)
}

func TestMetuxTwo(t *testing.T) {
	var mu sync.Mutex
	var count = 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Log(count)
}

func TestMetuxThree(t *testing.T) {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Log(counter.Count)
}

const (
	mutexLocked      = 1 << iota // mutex is locked //1
	mutexWoken                   //2
	mutexWaiterShift = iota      //2
)

func TestMetuxFour(t *testing.T) {
	t.Log(mutexLocked)
	t.Log(mutexWoken)
	t.Log(mutexWaiterShift)
}

func TestMetuxFive(t *testing.T) {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}

type Counter struct {
	sync.Mutex
	Count uint64
}

func TestMutexSix(t *testing.T) {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c, t) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter, t *testing.T) {
	c.Lock()
	defer c.Unlock()
	t.Log("in foo")
}
