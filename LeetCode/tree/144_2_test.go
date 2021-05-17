package tree

import "testing"

//二叉树的前序遍历 根 左 右

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(node *TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		helper(node.Left)
		helper(node.Right)

	}

	helper(root)

	return result
}

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 5}
	t.Log(preorderTraversal(root))
}
