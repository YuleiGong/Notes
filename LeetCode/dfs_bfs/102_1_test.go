package dfs_bfs

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//dfs
/*
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)

	return result
}
*/

//bfs
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root} //初始化队列

	for level := 0; 0 < len(queue); level++ { //每一层
		result = append(result, []int{})
		next := []*TreeNode{}
		for _, node := range queue {
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue = next
	}

	return result
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Left = &TreeNode{Val: 21}

	t.Logf("%v", levelOrder(root))
}
