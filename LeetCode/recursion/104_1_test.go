package recursion

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var helper func(*TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lcount := helper(root.Left)
		rcount := helper(root.Right)

		return max(lcount, rcount) + 1
	}

	return helper(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	t.Log(maxDepth(root))
}
