package linked_list

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}

	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return result.Next

}

func NewLinkedList(nums []int) (head *ListNode) {
	head = &ListNode{Val: nums[0], Next: nil}
	l := head
	for i, val := range nums {
		if i == 0 {
			continue
		}
		tmp := &ListNode{Val: val, Next: nil}
		head.Next = tmp
		head = head.Next
	}
	return l
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
func TestMergeTwoLists(t *testing.T) {
	l1 := NewLinkedList([]int{1, 2, 4})
	l2 := NewLinkedList([]int{1, 3, 4})
	printList(mergeTwoLists(l1, l2))
}
