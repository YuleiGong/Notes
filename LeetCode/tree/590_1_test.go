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
	var val []int
	var post func(*Node)
	post = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			post(child)
		}
		val = append(val, node.Val)
	}
	post(root)

	return val

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
