package recursion

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var helper func(root *TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lCount := helper(root.Left)
		rCount := helper(root.Right)

		//这部分容易忽略
		if root.Left == nil {
			return rCount + 1
		}
		if root.Right == nil {
			return lCount + 1
		}

		return min(lCount, rCount) + 1

	}

	return helper(root)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	t.Log(minDepth(root))

}
