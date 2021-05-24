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

func (this *Codec) serialize(root *TreeNode) string {
	var helper func(*TreeNode)
	var s []string

	helper = func(root *TreeNode) {
		if root == nil {
			s = append(s, "NULL")
			return
		}
		val := strconv.Itoa(root.Val)
		s = append(s, val)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return strings.Join(s, ",")
}

// Deserializes your encoded data to tree. 反序列化
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
	root = ser.deserialize(data)
	t.Log(root.Val)
}
