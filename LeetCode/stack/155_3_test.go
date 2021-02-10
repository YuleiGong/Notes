package stack

import (
	"testing"
)

type MinStack struct {
	stack    []int
	minStack []int
}

const INT_MAX = int(^uint(0) >> 1)

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{INT_MAX},
	}

}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, min(x, this.minStack[len(this.minStack)-1]))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
	this.minStack = this.minStack[0 : len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Pop()
	t.Log(s.GetMin())
}
