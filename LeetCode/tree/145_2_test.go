package tree

import "testing"

//二叉树的后序遍历 左 右 根

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(*TreeNode)

	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		result = append(result, node.Val)
	}
	helper(root)

	return result
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
