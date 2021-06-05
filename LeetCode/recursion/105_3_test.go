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

	var rootIndex int
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7} //前序 根左右
	inorder := []int{9, 3, 15, 20, 7}  //中序 左根右

	buildTree(preorder, inorder)
}
