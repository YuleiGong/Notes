package tree

import "testing"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	var helper func(root *Node)

	helper = func(root *Node) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		for _, node := range root.Children {
			helper(node)
		}
	}
	helper(root)

	return result
}

func TestPreorder(t *testing.T) {
	root := &Node{Val: 1}

	node1 := &Node{Val: 3}
	node1.Children = []*Node{&Node{Val: 5}, &Node{Val: 6}}

	node2 := &Node{Val: 2}
	node3 := &Node{Val: 4}
	root.Children = []*Node{node1, node2, node3}

	t.Log(preorder(root))
}
