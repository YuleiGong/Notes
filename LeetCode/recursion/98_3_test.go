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

//1:中序遍历 左 根 右
//2:比较大小
func isValidBST(root *TreeNode) bool {
	var last = math.MinInt64
	var helper func(root *TreeNode) bool

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
