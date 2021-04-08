# Channel
*** Don’t communicate by sharing memory, share memory by communicating.***
* 执行业务处理的 goroutine 不要通过共享内存的方式通信，而是要通过 Channel 通信的方式分享数据
* “communicate by sharing memory”是传统的并发编程处理方式，就是指，共享的数据需要用 __锁__ 进行保护，goroutine 需要获取到锁，才能并发访问数据。
* “share memory by communicating”则是类似于 CSP 模型的方式，通过通信的方式，一个 goroutine 可以把数据的“所有权”交给另外一个 goroutine


## 1 Channel 的应用场景
* 数据交流：当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个 goroutine 可以并发当作生产者（Producer）和消费者（Consumer）。
* 数据传递：一个 goroutine 将数据交给另一个 goroutine，相当于把数据的拥有权 (引用) 托付出去。
* 信号通知：一个 goroutine 可以将信号 (closing、closed、data ready 等) 传递给另一个或者另一组 goroutine 
* 任务编排：可以让一组 goroutine 按照一定的顺序并发或者串行的执行，这就是编排的功能。
* 锁：利用 Channel 也可以实现互斥锁的机制。

## 2 Channel 的基本用法
* chan 分为只能接收、只能发送、既可以接收又可以发送(双向)。三种类型:
```go
chan string          // 可以发送接收string 
chan<- struct{}      // 只能发送struct{}
<-chan int           // 只能从chan接收int
```
* 通过 make，我们可以初始化一个 chan，未初始化的 chan 的零值是 nil。你可以设置它的容量，比如下面的 chan 的容量是 9527，我们把这样的 chan 叫做 buffered chan；如果没有设置，它的容量是 0，我们把这样的 chan 叫做 unbuffered chan。  
```go
make(chan int, 9527)
```
* chan 中还有数据，那么，从这个 chan 接收数据的时候就不会阻塞，如果 chan 还未满（“满”指达到其容量），给它发送数据也不会阻塞，否则就会阻塞。unbuffered chan 只有__读写都准备好__之后才不会阻塞(边读边写)，这也是很多使用 unbuffered chan 时的常见 Bug。

* 发送数据 
```go
ch <- 2000
```
* 接收数据:
 * 接收数据时，还可以返回两个值。第一个值是返回的 chan 中的元素，很多人不太熟悉的是第二个值。第二个值是 bool 类型，代表是否成功地从 chan 中读取到一个值，如果第二个参数是 false，chan 已经被 __close__ 而且 chan 中没有缓存的数据，这个时候，第一个值是零值。所以，如果从 chan 读取到一个零值，可能是 sender 真正发送的零值，也可能是 closed 的并且没有缓存元素产生的零值。
 ```go
 x := <-ch // 把接收的一条数据赋值给变量x
 foo(<-ch) // 把接收的一个的数据作为参数传给函数
 <-ch // 丢弃接收的一条数据
 ```  
* 其他操作
 * Go 内建的函数 close、cap、len 都可以操作 chan 类型：close 会把 chan 关闭掉，cap 返回 chan 的容量，len 返回 chan 中缓存的还未被取走的元素数量。
 * chan 还可以应用于 for-range 语句中:
  ```go
  for v := range ch {
      fmt.Println(v)
  }
  ```
 * 忽略读取的值，只是清空 chan：
 ```go
 for range ch {
 }
 ```
     
## 3 channel 常见错误 
* close 为 nil 的 chan；
* send 已经 close 的 chan；
* close 已经 close 的 chan。
* goroutine 泄露:
 * 如果发生超时，process 函数就返回了，这就会导致 unbuffered 的 chan 从来就没有被读取。我们知道，unbuffered chan 必须等 reader 和 writer 都准备好了才能交流，否则就会阻塞。超时导致未读，结果就是子 goroutine 就阻塞在第 7 行永远结束不了，进而导致 goroutine 泄漏。将 unbuffered chan 改成容量为 1 的 chan，这样第 7 行就不会被阻塞了。

 ```go
 func process(timeout time.Duration) bool {
     ch := make(chan bool)
 
     go func() {
         // 模拟处理耗时的业务
         time.Sleep((timeout + time.Second))
         ch <- true // block
         fmt.Println("exit goroutine")
     }()
     select {
     case result := <-ch:
         return result
     case <-time.After(timeout):
         return false
     }
 }
 ```
