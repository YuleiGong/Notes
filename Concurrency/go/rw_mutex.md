# RWMutex 

* 如果某个读操作的 goroutine 持有了锁，在这种情况下，其它读操作的 goroutine 就不必一直傻傻地等待了，而是可以并发地访问共享变量，这样我们就可以将串行的读变成并行读，提高读操作的性能。当写操作的 goroutine 持有锁的时候，它就是一个排外锁，其它的写操作和读操作的 goroutine，需要阻塞等待持有这个锁的 goroutine 释放锁。

## 1 读写锁的使用

* __Lock/Unlock__ ：写操作时调用的方法。如果锁已经被 reader 或者 writer 持有，那么，Lock 方法会一直阻塞，直到能获取到锁；Unlock 则是配对的释放锁的方法。
* __RLock/RUnlock__ ：读操作时调用的方法。如果锁已经被 writer 持有的话，RLock 方法会一直阻塞，直到能获取到锁，否则就直接返回；而 RUnlock 是 reader 释放锁的方法。
* 如果你遇到可以明确区分 reader 和 writer goroutine 的场景，且有 __大量的并发读、少量的并发写__，并且有强烈的性能需求，你就可以考虑使用读写锁 RWMutex 替换 Mutex。

```go
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
```

## 2 实现原理

* Go 标准库中的 RWMutex 是基于 Mutex 实现的。
* __Read-preferring__：读优先的设计可以提供很高的并发性，但是，在竞争激烈的情况下可能会导致写饥饿。这是因为，如果有大量的读，这种设计会导致只有所有的读都释放了锁之后，写才可能获取到锁。
* __Write-preferring__：写优先的设计意味着，如果已经有一个 writer 在等待请求锁的话，它会 __阻止__ 新来的请求锁的 reader 获取到锁，所以优先保障 writer。当然，如果有一些 reader 已经请求了锁的话，新请求的 writer 也会等待已经存在的 reader 都释放锁之后才能获取。所以，写优先级设计中的优先权是针对新来的请求而言的。这种设计主要避免了 writer 的饥饿问题。
* __不指定优先级__：这种设计比较简单，不区分 reader 和 writer 优先级，某些场景下这种不指定优先级的设计反而更有效，因为第一类优先级会导致写饥饿，第二类优先级可能会导致读饥饿，这种不指定优先级的访问不再区分读写，大家都是同一个优先级，解决了饥饿的问题。

***Go 标准库中的 RWMutex 设计是 Write-preferring 方案。一个正在阻塞的 Lock 调用会排除新的 reader 请求到锁。***

## 3 RWMutex 错误使用场景

### 3.1 不可复制

* RWMutex 是由一个互斥锁和四个辅助字段组成的。我们很容易想到，互斥锁是不可复制的，再加上四个有状态的字段，RWMutex 就更加不能复制使用了。 

### 3.1 不可重入

* 1 读写锁因为重入（或递归调用）导致死锁: 因为读写锁内部基于互斥锁实现对 writer 的并发访问，而互斥锁本身是有重入问题的，所以，writer 重入调用 Lock 的时候，就会出现死锁的现象
* 2 有活跃 reader 的时候，writer 会等待，如果我们在 reader 的读操作时调用 writer 的写操作（它会调用 Lock 方法），那么，这个 reader 和 writer 就会形成互相依赖的死锁状态。Reader 想等待 writer 完成后再释放锁，而 writer 需要这个 reader 释放锁之后，才能不阻塞地继续执行。这是一个读写锁常见的死锁场景。
* 3 当一个 writer 请求锁的时候，如果已经有一些活跃的 reader，它会等待这些活跃的 reader 完成，才有可能获取到锁，但是，如果之后活跃的 reader 再依赖新的 reader 的话，这些新的 reader 就会等待 writer 释放锁之后才能继续执行，这就形成了一个环形依赖： writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader 依赖 writer。

![rwlock.jpg](https://i.loli.net/2021/04/05/Tul4AF1YMjIkxw7.jpg)

### 3.2 释放未加锁的 RWMutex

* 释放为加锁的RWMutex 会导致Panic
