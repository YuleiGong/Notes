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
		stack:    make([]int, 0),
		minStack: []int{math.MaxInt8},
	}

}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.minStack = append(this.minStack, min(x, this.minStack[len(this.minStack)-1]))
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
	s.Pop()
	t.Log(s.GetMin())
}
