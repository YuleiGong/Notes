package stack

import (
	"math"
	"testing"
)

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}

}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1] //最小
	this.minStack = append(this.minStack, min(top, x))

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]

}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	t.Log(s.GetMin())
}
