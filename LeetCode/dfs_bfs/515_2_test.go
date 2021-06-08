package dfs_bfs

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for level := 0; 0 < len(queue); level++ {
		next := []*TreeNode{}

		max := math.MinInt64
		for _, node := range queue {
			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

		}
		result = append(result, max)
		queue = next
	}

	return result
}

func TestLargestValues(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 9}

	t.Logf("%v", largestValues(root))
}
