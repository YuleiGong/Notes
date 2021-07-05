package recursion

import (
	"strconv"
	"strings"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Seriaizes a tree to a single string. 前序 根左右
func (this *Codec) serialize(root *TreeNode) string {
	var result []string
	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			result = append(result, "NULL")
			return
		}

		result = append(result, strconv.Itoa(root.Val))
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree. 根左右
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	var helper func() *TreeNode

	helper = func() *TreeNode {
		if s[0] == "NULL" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		root := &TreeNode{Val: val}
		s = s[1:]
		root.Left = helper()
		root.Right = helper()

		return root
	}

	return helper()
}

func TestSerializeOrDeserialize(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	ser := Constructor()
	data := ser.serialize(root)
	t.Log(data)
	root = ser.deserialize(data)
	t.Log(root.Val)
}
