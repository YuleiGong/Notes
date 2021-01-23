package linked_list

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for slow != nil {
				if p == slow {
					return p
				}

				p = p.Next
				slow = slow.Next
			}
		}
	}

	return nil
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
