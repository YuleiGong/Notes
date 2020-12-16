package linked_list

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next

	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
	return false

}

func TestHasCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next = head.Next
	t.Log(hasCycle(head))
}
