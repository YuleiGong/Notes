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
	var helper func(*TreeNode) bool
	var min = math.MinInt64

	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !helper(root.Left) {
			return false
		}

		if root.Val <= min {
			return false
		} else {
			min = root.Val
		}

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

	/*
		root := &TreeNode{
			Val: 5,
		}
		root.Left = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 4}
		root.Right.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 6}
	*/

	t.Log(isValidBST(root))
}
