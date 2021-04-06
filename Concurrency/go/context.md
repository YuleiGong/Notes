# Context

* 上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息
* 控制子 goroutine 的运行
* 超时控制的方法调用
* 可以取消的方法调用

## 1. Context 的基本使用

* Context 中实现了 2 个常用的生成顶层 Context 的方法: context.Background(), context.TODO(), 可以直接使用 context.Background。事实上，它们两个底层的实现是一模一样的
* Context 包含如下几个方法:
  * Done() 方法返回一个 Channel 对象。在 Context 被取消时，此 Channel 会被 close，如果没被取消，可能会返回 nil。后续的 Done 调用总是返回相同的结果。当 Done 被 close 的时候，你可以通过 ctx.Err 获取错误信息。如果 Done 没有被 close，Err 方法返回 nil；如果 Done 被 close，Err 方法会返回 Done 被 close 的原因。
  * Value() 返回此 ctx 中和指定的 key 相关联的 value。
  * Deadlien() 方法会返回这个 Context 被取消的截止日期。如果没有设置截止日期，ok 的值是 false。后续每次调用这个对象的 Deadline 方法时，都会返回和第一次调用相同的结果。

### 1.1 WithValue

* WithValue 基于 __parent Context__ 生成一个新的 Context，保存了一个 key-value 键值对。它常常用来传递上下文。
* WithValue 方法其实是创建了一个类型为 valueCtx 的 Context，它的类型定义如下：

```go
type valueCtx struct {
    Context
    key, val interface{}
}
```

* Go 标准库实现的 Context 还实现了链式查找。如果不存在，还会向 parent Context 去查找。
![ctx_value.jpg](https://i.loli.net/2021/04/06/nAOW1mwKdp9HU8q.jpg)

```go
func main() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0001")
	ctx = context.WithValue(ctx, "key3", "0001")
	ctx = context.WithValue(ctx, "key4", "0004")

	fmt.Println(ctx.Value("key1"))
}
```

### 1.2 WithCancel

* WithCancel 方法返回 parent 的副本，它的类型是 cancelCtx。
* 我们常常在一些需要主动取消长时间的任务时，创建这种类型的 Context，然后把这个 Context 传给长时间执行任务的 goroutine。当需要中止任务时，我们就可以执行这个cancel 函数，这样长时间执行任务的 goroutine，就可以依次结束。
* cancel 是向下传递的，如果一个 WithCancel 生成的 Context 被 cancel 时，如果它的子 Context（也有可能是孙，或者更低，依赖子的类型）也是 cancelCtx 类型的，就会被 cancel，但是不会向上传递。parent Context 不会因为子 Context 被 cancel 而 cancel。

```go
func main() {
	// 父context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go watch1(ctx)

	// 子context
	valueCtx, _ := context.WithCancel(ctx)
	go watch2(valueCtx)

	time.Sleep(3 * time.Second)


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
```
### 1.3 WithTimeOut

* WithTimeout 其实是和 WithDeadline 一样，只不过一个参数是超时时间，一个参数是截止时间。超时时间加上当前时间，其实就是截止时间，因此，WithTimeout 的实现是：

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    // 当前时间+timeout就是deadline
    return WithDeadline(parent, time.Now().Add(timeout))
}
```

* 只要你的任务正常完成了，就需要调用 cancel，这样，这个 Context 才能释放它的资源

```go
func main() {
	// 创建一个子节点的context,3秒后自动超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go watch(ctx, "watch1")
	go watch(ctx, "watch2")

	time.Sleep(8 * time.Second)

	fmt.Println("cancel")

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
```

### 1.4 WithDeadline

* WithDeadline 会返回一个 parent 的副本，并且设置了一个不晚于参数 d 的截止时间，类型为 timerCtx。
* 如果它的截止时间晚于 parent 的截止时间，那么就以 parent 的截止时间为准，并返回一个类型为 cancelCtx 的 Context，因为 parent 的截止时间到了，就会取消这个 cancelCtx。
* 如果当前时间已经超过了截止时间，就直接返回一个已经被 cancel 的 timerCtx。否则就会启动一个定时器，到截止时间取消这个 timerCtx。

```go
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()

	go watch(ctx, "watch1")
	go watch(ctx, "watch2")

	time.Sleep(8 * time.Second)

	fmt.Println("cancel")

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
```