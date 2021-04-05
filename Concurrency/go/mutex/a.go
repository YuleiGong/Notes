	func (m *Mutex) Unlock() {
		// Fast path: drop lock bit.
		new := atomic.AddInt32(&m.state, -mutexLocked) //去掉锁标志
		if (new+mutexLocked)&mutexLocked == 0 {        //本来就没有加锁
			panic("sync: unlock of unlocked mutex")
		}

		old := new
		for {
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 { // 没有等待者，或者有唤醒的waiter，或者锁原来已加锁
				return
			}
			new = (old - 1<<mutexWaiterShift) | mutexWoken // 新状态，准备唤醒goroutine，并设置唤醒标志
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime.Semrelease(&m.sema)
				return
			}
			old = m.state
		}
	}
