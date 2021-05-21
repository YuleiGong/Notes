package tree

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历顺序 根 左 右
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var helper func(*TreeNode)

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
