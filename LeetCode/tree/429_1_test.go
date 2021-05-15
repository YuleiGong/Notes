package tree

//N 叉树层序遍历

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var val [][]int

	if root == nil {
		return val
	}

	var helper func(*Node, int)
	helper = func(root *Node, level int) {
		if root == nil {
			return
		}
		if len(val) <= level {
			val = append(val, []int{})
		}
		val[level] = append(val[level], root.Val)
		for _, c := range root.Children {
			helper(c, level+1)
		}
	}
	helper(root, 0)

	return val

}
func TestLevelOrder(t *testing.T) {
	root := &Node{Val: 1}

	node1 := &Node{Val: 3}
	node1.Children = []*Node{&Node{Val: 5}, &Node{Val: 6}}

	node2 := &Node{Val: 2}
	node3 := &Node{Val: 4}
	root.Children = []*Node{node1, node2, node3}

	t.Log(levelOrder(root))
}
