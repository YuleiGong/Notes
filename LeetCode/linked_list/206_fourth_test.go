package linked_list

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev

}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func TestReverseList(t *testing.T) {
	cur := &ListNode{}
	head := cur
	for _, val := range []int{1, 2, 3, 4, 5} {
		cur.Val = val
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	printList(reverseList(head))

}
