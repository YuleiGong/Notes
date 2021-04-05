# WaitGroup

## 1 基本用法

* __Add__ ，用来设置 WaitGroup 的计数值；
* __Done__ ，用来将 WaitGroup 的计数值减 1，其实就是调用了 Add(-1)；
* __Wait__ ，调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0。

```go
// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取当前的计数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// sleep 1秒，然后计数值加1
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10) // WaitGroup的值设置为10

	for i := 0; i < 10; i++ { // 启动10个goroutine执行加1任务
		go worker(&counter, &wg) //wg 
	}
	// 检查点，等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}
```

## 2 WaitGroup 错误使用场景

### 2.1 计数器设置为负值 

* 调用 Add 的时候传递一个负数。如果你能保证当前的计数器加上这个负数后还是大于等于 0 的话，也没有问题，否则就会导致 panic。

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	wg.Add(-10) //将-10作为参数调用Add，计数值被设置为0
	wg.Add(-1) //将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发panic
}
```

### 2.2 调用 Done 方法的次数过多，超过了 WaitGroup 的计数值。

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Done()
}
```

### 2.3 不期望的 Add 时机

* 在使用 WaitGroup 的时候，你一定要遵循的原则就是，__等所有的 Add 方法调用之后再调用 Wait，否则就可能导致 panic 或者不期望的结果__。


### 2.4 前一个 Wait 还没结束就重用 WaitGroup

* WaitGroup 是可以重用的。只要 WaitGroup 的计数值恢复到零值的状态，那么它就可以被看作是新创建的 WaitGroup，被重复使用。
* 如果我们在 WaitGroup 的计数值还没有恢复到零值的时候就重用，就会导致程序 panic,WaitGroup 虽然可以重用，但是是有一个前提的，那就是必须等到上一轮的 Wait 完成之后，才能重用 WaitGroup 执行下一轮的 Add/Wait，如果你在 Wait 还没执行完的时候就调用下一轮 Add 方法，就有可能出现 panic。

### 2.5 copy WaitGroup 

```go
type TestStruct struct {
  Wait sync.WaitGroup
}

func main() {
  w := sync.WaitGroup{}
  w.Add(1)
  t := &TestStruct{
    Wait: w, //copy 了WaitGroup 的实例w
  }

  t.Wait.Done()
  fmt.Println("Finished")
}
```