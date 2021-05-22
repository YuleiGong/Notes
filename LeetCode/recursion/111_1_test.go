package recursion

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var helper func(*TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lcount := helper(root.Left)
		rcount := helper(root.Right)
		if root.Left == nil {
			return rcount + 1
		}
		if root.Right == nil {
			return lcount + 1
		}
		return min(rcount, lcount) + 1
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
