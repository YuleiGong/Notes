# Mutex

* __互斥锁__: 互斥锁是并发控制的一个基本手段，是为了避免竞争而建立的一种并发控制机制。
* __临界区__: 在并发编程中，如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，这部分被保护起来的程序，就叫做 __临界区__。如果很多线程同步访问临界区，就会造成访问或操作错误，这当然不是我们希望看到的结果。所以，我们可以使用互斥锁，限定临界区只能同时由一个线程持有。当临界区由一个线程持有的时候，其它线程如果想进入这个临界区，就会返回失败，或者是等待。直到持有的线程退出临界区，这些等待线程中的某一个才有机会接着持有这个临界区。
![获取锁.jpg](https://i.loli.net/2021/04/01/Rj83gsfeLOZuS4p.jpg)
* Mutex 是使用最广泛的 __同步原语__  

## 1. Mutex 的基本使用

* 互斥锁 Mutex 就提供两个方法 Lock 和 Unlock：进入临界区之前调用 Lock 方法，退出临界区的时候调用 Unlock 方法
* 当一个 goroutine 通过调用 Lock 方法获得了这个锁的拥有权后， 其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，直到锁被释放并且自己获取到了这个锁的拥有权。

* 示例1:
  * count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。

  ```go
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
  ``` 

  * 通过 race 参数，可以在测试初期发现并发问题。

   ```go
   ➜  go git:(master) ✗ go test mutex_test.go -v -race
   === RUN   TestMetuxOne
   ==================
   WARNING: DATA RACE
   Read at 0x00c000016128 by goroutine 9:
     command-line-arguments.TestMetuxOne.func1()
         /Users/gongyulei/Documents/my_gitbook/Notes/Concurrency/go/mutex_test.go:16 +0x78
  
   Previous write at 0x00c000016128 by goroutine 8:
     command-line-arguments.TestMetuxOne.func1()
         /Users/gongyulei/Documents/my_gitbook/Notes/Concurrency/go/mutex_test.go:16 +0x91
  
   Goroutine 9 (running) created at:
     command-line-arguments.TestMetuxOne()
         /Users/gongyulei/Documents/my_gitbook/Notes/Concurrency/go/mutex_test.go:13 +0xe4
     testing.tRunner()
         /usr/local/opt/go/libexec/src/testing/testing.go:991 +0x1eb
  
   Goroutine 8 (running) created at:
     command-line-arguments.TestMetuxOne()
         /Users/gongyulei/Documents/my_gitbook/Notes/Concurrency/go/mutex_test.go:13 +0xe4
     testing.tRunner()
         /usr/local/opt/go/libexec/src/testing/testing.go:991 +0x1eb
   ==================
       TestMetuxOne: mutex_test.go:21: 212385
       TestMetuxOne: testing.go:906: race detected during execution of test
   --- FAIL: TestMetuxOne (0.40s)
       : testing.go:906: race detected during execution of test
   FAIL
   FAIL    command-line-arguments  0.760s
   FAIL
   ```

* 示例2
  * Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）

  ```go
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
  ```

  * 还可以采用嵌入字段的方式。通过嵌入字段，你可以在这个 struct 上直接调用 Lock/Unlock 方法。 

  ```go
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

  type Counter struct {
  	sync.Mutex
  	Count uint64
  }
  ```

## 2. Mutex 实现  

### 2.1 初版互斥锁实现

* __CAS__ 指令将给定的值和一个内存地址中的值进行比较，如果它们是同一个值，就使用新值替换内存地址中的值，这个操作是原子性的。
![lock.jpg](https://i.loli.net/2021/04/04/6Z2sKbgtxrDkYnA.jpg)
* __Lock__: 调用 Lock 请求锁的时候，通过 xadd 方法进行 CAS 操作，xadd 方法通过循环执行CAS操作直到成功，保证对key 加1的操作成功完成。如果比较幸运，锁没有被别的 goroutine 持有，那么，Lock 方法成功地将 key 设置为 1，这个 goroutine 就持有了这个锁；如果锁已经被别的 goroutine 持有了，那么，当前的 goroutine 会把 key 加 1，而且还会调用 semacquire 方法，使用信号量将自己休眠，等锁释放的时候，信号量会将它唤醒。
* __Unlock__ 持有锁的 goroutine 调用 Unlock 释放锁时，它会将 key 减 1。如果当前没有其它等待这个锁的 goroutine，这个方法就返回了。但是，如果还有等待此锁的其它 goroutine，那么，它会调用 semrelease 方，利用信号量唤醒等待锁的其它 goroutine 中的一个。
* __Unlock__ 方法可以被任意的 goroutine 调用释放锁，即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。这是因为，Mutex 本身并没有包含持有这把锁的 goroutine 的信息，所以，Unlock 也不会对此进行检查。Mutex 的这个设计一直保持至今。其它 goroutine 可以强制释放锁，这是一个非常危险的操作，因为在临界区的 goroutine 可能不知道锁已经被释放了，还会继续执行临界区的业务操作，这可能会带来意想不到的结果。一定要遵循 __谁申请，谁释放__ 的原则。在真实的实践中，我们使用互斥锁的时候，很少在一个方法中单独申请锁，而在另外一个方法中单独释放锁，一般都会在同一个方法中获取锁和释放锁。

```go
   // CAS操作，当时还没有抽象出atomic包
    func cas(val *int32, old, new int32) bool
    func semacquire(*int32)
    func semrelease(*int32)
    // 互斥锁的结构，包含两个字段
    type Mutex struct {
        key  int32 // 锁是否被持有的标
        sema int32 // 信号量专用，用以阻塞/唤醒goroutine
    }
    
    // 保证成功在val上增加delta的值
    func xadd(val *int32, delta int32) (new int32) {
        for {
            v := *val
            if cas(val, v, v+delta) {
                return v + delta
            }
        }
        panic("unreached")
    }
    
    // 请求锁
    func (m *Mutex) Lock() {
        if xadd(&m.key, 1) == 1 { //标识加1，如果等于1，成功获取到锁
            return
        }
        semacquire(&m.sema) // 否则阻塞等待
    }
    
    func (m *Mutex) Unlock() {
        if xadd(&m.key, -1) == 0 { // 将标识减去1，如果等于0，则没有其它等待者
            return
        }
        semrelease(&m.sema) // 唤醒其它阻塞的goroutine
    }    
```

## 3. Mutex 错误使用场景

### 3.1 Lock/Unlock 不是成对出现

* Lock/Unlock 没有成对出现，就意味着会出现死锁的情况(没有Unlock 其他的goroutine 获取不到锁)，或者是因为 Unlock 一个未加锁的 Mutex 而导致 panic。

```go
func TestMetuxFive(t *testing.T) {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}
out:
=== RUN   TestMetuxFive
hello world!
fatal error: sync: unlock of unlocked mutex
```

### 3.2 Copy 已使用的 Mutex

* Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。如果你要复制一个已经加锁的 Mutex 给一个新的变量，那么新的刚初始化的变量居然被加锁了，这显然不符合你的期望，因为你期望的是一个零值的 Mutex。关键是在并发环境下，你根本不知道要复制的 Mutex 状态是什么，因为要复制的 Mutex 是由其它 goroutine 并发访问的，状态可能总是在变化。

```go
type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
```

* 复制之前已经使用了这个锁，这就导致，复制的 Counter 是一个带状态 Counter。foo()函数一直无法获取锁。
* 使用 vet 工具可以帮助及早发现有copy Mutes的情况:

```go
➜  go git:(master) ✗ go vet vet.go
# command-line-arguments
./vet.go:18:6: call of foo copies lock value: command-line-arguments.Counter
./vet.go:22:12: foo passes lock by value: command-line-arguments.Counter
```

### 3.3 重入

* 当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果 __拥有这把锁的线程__ 再请求这把锁的话，不会阻塞，而是成功返回，所以叫可重入锁（有时候也叫做递归锁）。只要你拥有这把锁，你可以可着劲儿地调用，比如通过递归实现一些算法，调用者不会阻塞或者死锁。
* 可重入锁（递归锁）解决了代码重入或者递归调用带来的死锁问题，同时它也带来了另一个好处，就是我们可以要求，只有持有锁的 goroutine 才能 unlock 这个锁。
* Mutex __不是可重入的锁__ 。Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何 goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件。

```
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

out:
➜  go git:(master) ✗ go run repeat_lock.go
in foo
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_SemacquireMutex(0xc0000140b4, 0x1175e00, 0x1)
        /usr/local/opt/go/libexec/src/runtime/sema.go:71 +0x47
sync.(*Mutex).lockSlow(0xc0000140b0)
        /usr/local/opt/go/libexec/src/sync/mutex.go:138 +0xfc
sync.(*Mutex).Lock(0xc0000140b0)
        /usr/local/opt/go/libexec/src/sync/mutex.go:81 +0x47
main.bar(0x10ea620, 0xc0000140b0)
        /Users/gongyulei/Documents/Notes/Concurrency/go/repeat_lock.go:16 +0x35
main.foo(0x10ea620, 0xc0000140b0)
        /Users/gongyulei/Documents/Notes/Concurrency/go/repeat_lock.go:11 +0xa5
main.main()
        /Users/gongyulei/Documents/Notes/Concurrency/go/repeat_lock.go:23 +0x3d
exit status 2
```

#### 3.3.1 使用goroutine id 设计可重入锁

* 我们用 owner 字段，记录当前锁的拥有者 goroutine 的 id；recursion 是辅助字段，用于记录重入的次数
* 尽管拥有者可以多次调用 Lock，但是也必须调用相同次数的 Unlock，这样才能把锁释放掉。

```go
// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
    sync.Mutex
    owner     int64 // 当前持有锁的goroutine id
    recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
    gid := goid.Get()
    // 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
    if atomic.LoadInt64(&m.owner) == gid {
        m.recursion++
        return
    }

    m.Mutex.Lock()
    // 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
    atomic.StoreInt64(&m.owner, gid)
    m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
    gid := goid.Get()
    // 非持有锁的goroutine尝试释放锁，错误的使用
    if atomic.LoadInt64(&m.owner) != gid {
        panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
    }
    // 调用次数减1
    m.recursion--
    if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
        return
    }
    // 此goroutine最后一次调用，需要释放锁
    atomic.StoreInt64(&m.owner, -1)
    m.Mutex.Unlock()
}
```

#### 3.3.2 使用token 设计可重入锁

* Go 开发者不期望你利用 goroutine id 做一些不确定的东西，所以，他们没有暴露获取 goroutine id 的方法。
* 调用者自己提供一个 token，获取锁的时候把这个 token 传入，释放锁的时候也需要把这个 token 传入。通过用户传入的 token 替换goroutine id，

```go
// Token方式的递归锁
type TokenRecursiveMutex struct {
    sync.Mutex
    token     int64
    recursion int32
}

// 请求锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
    if atomic.LoadInt64(&m.token) == token { //如果传入的token和持有锁的token一致，说明是递归调用
        m.recursion++
        return
    }
    m.Mutex.Lock() // 传入的token不一致，说明不是递归调用
    // 抢到锁之后记录这个token
    atomic.StoreInt64(&m.token, token)
    m.recursion = 1
}

// 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
    if atomic.LoadInt64(&m.token) != token { // 释放其它token持有的锁
        panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
    }
    m.recursion-- // 当前持有这个锁的token释放锁
    if m.recursion != 0 { // 还没有回退到最初的递归调用
        return
    }
    atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
    m.Mutex.Unlock()
}
```

### 3.4 死锁

* 两个或两个以上的进程（或线程，goroutine）在执行过程中，因争夺共享资源而处于一种互相等待的状态，如果没有外部干涉，它们都将无法推进下去，此时，我们称系统处于死锁状态或系统产生了 __死锁__。


## 4 扩展Mutex

### 4.1 实现 TryLock

* 当一个 goroutine 调用这个 TryLock 方法请求锁的时候，如果这把锁没有被其他 goroutine 所持有，那么，这个 goroutine 就持有了这把锁，并返回 true；如果这把锁已经被其他 goroutine 所持有，或者是正在准备交给某个被唤醒的 goroutine，那么，这个请求锁的 goroutine 就直接返回 false，__不会阻塞__ 在方法调用上。

```go
// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// 扩展一个Mutex结构
type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁,如果幸运，没有其他 goroutine 争这把锁，那么，这把锁就会被这个请求的 goroutine 获取，直接返回。
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁。这个时候，可能还有其他的 goroutine 也在竞争这把锁，所以，不能保证成功获取这把锁。
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
```

### 4.2 实现线程安全队列

* Mutex 经常会和其他非线程安全（对于 Go 来说，我们其实指的是 goroutine 安全）的数据结构一起，组合成一个线程安全的数据结构。新数据结构的业务逻辑由原来的数据结构提供，而 Mutex 提供了锁的机制，来保证线程安全。

```go
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
```
