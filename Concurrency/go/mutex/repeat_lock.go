package main

import (
	"fmt"
	"sync"
)

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) { //同一个goroutine 再次拿锁。因为不可重入，会死锁
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	l := &sync.Mutex{}
	foo(l)
}
