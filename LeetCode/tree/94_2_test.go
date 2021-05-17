package tree

import "testing"

//二叉树的中序遍历 左 根 右

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(*TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		result = append(result, node.Val)
		helper(node.Right)
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
