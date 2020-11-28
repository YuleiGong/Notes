package queue

type MyCircularQueue struct {
	Queue    []int //队列数组
	Head     int   //头指针
	Count    int   //队列实际大小
	Capacity int   //队列总长度
}

//初始化
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Queue:    make([]int, k, k),
		Head:     0, //队列头
		Count:    0, //队列中的元素个数
		Capacity: k, //队列的容量
	}
}

//入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Count >= this.Capacity {
		return false
	}
	this.Queue[(this.Head+this.Count)%this.Capacity] = value
	this.Count++
	return true
}

//出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.Count == 0 {
		return false
	}
	//删数组的头
	this.Head = this.Head + 1%this.Capacity
	this.Count--
	return true
}

//队头
func (this *MyCircularQueue) Front() int {
	if this.Count == 0 {
		return -1
	}
	return this.Queue[this.Head]
}

//队尾
func (this *MyCircularQueue) Rear() int {
	if this.Count == 0 {
		return -1
	}
	return this.Queue[(this.Head+this.Count-1)%this.Capacity]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Count == this.Capacity
}
