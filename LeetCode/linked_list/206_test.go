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
	var prev *ListNode = nil //前驱指针
	for head != nil {
		next := head.Next //记录下一个指针的位置，防止丢失
		head.Next = prev  //修改指针的位置
		prev = head       //修改前驱指针的位置为当前指针
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
