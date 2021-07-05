package recursion

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var rootIndex int //中序遍历中根节点的位置
	for index, val := range inorder {
		if val == preorder[0] {
			rootIndex = index
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7} //前序 根左右
	inorder := []int{9, 3, 15, 20, 7}  //中序 左根右

	buildTree(preorder, inorder)
}
