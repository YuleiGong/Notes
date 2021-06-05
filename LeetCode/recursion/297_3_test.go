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

// Serializes a tree to a single string. 前序
func (this *Codec) serialize(root *TreeNode) string {
	var result []string
	var helper func(*TreeNode)

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

// Deserializes your encoded data to tree. 还原 左右根
func (this *Codec) deserialize(data string) *TreeNode {
	result := strings.Split(data, ",")
	var helper func() *TreeNode

	helper = func() *TreeNode {
		if result[0] == "NULL" {
			result = result[1:]
			return nil
		}
		val, _ := strconv.Atoi(result[0])
		root := &TreeNode{Val: val}
		result = result[1:]
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
