package tree

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)
		result = append(result, root.Val)
		helper(root.Right)
	}
	helper(root)

	return result
}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 5}
	t.Log(inorderTraversal(root))
}
