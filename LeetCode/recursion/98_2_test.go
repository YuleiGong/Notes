package recursion

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var helper func(*TreeNode) bool
	var last = math.MinInt64

	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !helper(root.Left) || root.Val <= last {
			return false
		}
		last = root.Val

		return helper(root.Right)
	}

	return helper(root)
}

func TestIsValudBST(t *testing.T) {
	root := &TreeNode{
		Val: 2,
	}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	t.Log(isValidBST(root))
}
