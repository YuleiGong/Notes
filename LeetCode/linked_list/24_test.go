package linked_list

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head} //创建哑节点
	temp := dummyHead

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func TestSwapPairs(t *testing.T) {
	cur := &ListNode{}
	head := cur
	for _, val := range []int{1, 2, 3, 4} {
		cur.Val = val
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	swapPairs(head)
}
