package tree

import "testing"

//二叉树的后序遍历 左 右 根

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var vals []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		vals = append(vals, node.Val)
	}
	postorder(root)

	return vals
}

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 5}
	t.Log(postorderTraversal(root))
}
