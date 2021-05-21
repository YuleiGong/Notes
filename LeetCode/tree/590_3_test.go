package tree

import "testing"

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var result []int
	var helper func(*Node)

	helper = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			helper(child)
		}
		result = append(result, node.Val)
	}

	helper(root)

	return result
}

func TestPostorder(t *testing.T) {
	root := &Node{Val: 1}

	node1 := &Node{Val: 3}
	node1.Children = []*Node{&Node{Val: 5}, &Node{Val: 6}}

	node2 := &Node{Val: 2}
	node3 := &Node{Val: 4}
	root.Children = []*Node{node1, node2, node3}

	t.Log(postorder(root))
}
