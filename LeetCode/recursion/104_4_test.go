package recursion

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var helper func(*TreeNode, int) int

	helper = func(root *TreeNode, level int) int {
		if root == nil {
			return level
		}

		lCount := helper(root.Left, level+1)
		rCount := helper(root.Right, level+1)

		return max(rCount, lCount)

	}

	return helper(root, 0)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	//root.Right.Left = &TreeNode{Val: 15}
	//root.Right.Right = &TreeNode{Val: 7}

	t.Log(maxDepth(root))

}
