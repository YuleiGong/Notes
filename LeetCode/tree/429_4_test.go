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
	var result [][]int
	var helper func(root *Node, level int)

	helper = func(root *Node, level int) {
		if root == nil {
			return
		}

		if len(result) == level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		for _, node := range root.Children {
			helper(node, level+1)
		}
	}
	helper(root, 0)

	return result
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
