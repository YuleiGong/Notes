package linked_list

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//hash-map
/*
func detectCycle(head *ListNode) *ListNode {
	exist := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := exist[head]; ok {
			return head
		}
		exist[head] = true
		head = head.Next
	}

	return nil
}
*/
//快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slow *ListNode
	var ok bool
	if slow, ok = hasCycle(head); !ok {
		return nil
	}

	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

//首先判定是否有环
func hasCycle(head *ListNode) (*ListNode, bool) {
	fast := head
	slow := head

	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return slow, true
		}
	}

	return nil, false
}

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	node := detectCycle(head)
	t.Log(node.Val)
}
