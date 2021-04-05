# Once

* Once 常常用来初始化 __单例资源__，或者并发访问只需初始化一次的共享资源，或者在测试的时候初始化一次测试资源。

## 1. Once 的基本使用

* sync.Once 只暴露了一个方法 Do，你可以多次调用 Do 方法，但是只有第一次调用 Do 方法时 f 参数才会执行，这里的 f 是一个无参数无返回值的函数。

```go
func main() {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1) // 打印出 in f1

	// 第二个初始化函数
	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2) // 无输出
}
```

* f 参数是一个无参数无返回的函数，所以你可能会通过闭包的方式引用外面的参数。

```go
var addr = "baidu.com"

var conn net.Conn
var err error

once.Do(func() {
    conn, err = net.Dial("tcp", addr)
})
```

## 2. 错误使用场景

### 2.1 死锁

* Do 方法会执行一次 f，但是如果 f 中再次调用这个 Once 的 Do 方法的话，就会导致死锁的情况出现

```go
func main() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("初始化")
		})
	})
}
```

### 2.2 f()函数未初始化

* 如果 f 方法执行的时候 panic，或者 f 执行初始化资源的时候失败了，这个时候，Once 还是会认为初次执行已经成功了，即使再次调用 Do 方法，也不会再次执行 f。
* 自己实现一个类似 Once 的并发原语，既可以返回当前调用 Do 方法是否正确完成，还可以在初始化失败后调用 Do 方法再次尝试初始化，直到初始化成功才不再初始化了。

```go
type Once struct {
	m    sync.Mutex
	done uint32
}

// 传入的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do方法会把这个error返回给调用者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 { //fast path
		return nil
	}
	return o.slowDo(f)
}

// 如果还没有初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双检查，还没有初始化
		err = f()
		if err == nil { // 初始化成功才将标记置为已初始化
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
```
