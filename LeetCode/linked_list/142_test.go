package linked_list

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	exist := map[*ListNode]bool{}
	for head != nil {
		if _, ok := exist[head]; ok {
			return head
		}
		exist[head] = true
		head = head.Next
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next = head.Next

	node := detectCycle(head)
	t.Log(node.Val)
}
