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

// Serializes a tree to a single string. 序列化 使用前序遍历序列化
func (this *Codec) serialize(root *TreeNode) string {
	var s []string
	var helper func(*TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			s = append(s, "null")
			return
		}
		s = append(s, strconv.Itoa(root.Val))
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return strings.Join(s, ",")

}

// Deserializes your encoded data to tree. 反序列化
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	var helper func() *TreeNode

	helper = func() *TreeNode {
		if l[0] == "null" {
			l = l[1:]
			return nil
		}
		val, _ := strconv.Atoi(l[0])
		root := &TreeNode{Val: val}
		l = l[1:]
		root.Left = helper()
		root.Right = helper()

		return root
	}

	return helper()
}

func TestEncodeOrDecode(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	ser := Constructor()
	t.Log(ser.serialize(root))
	droot := ser.deserialize(ser.serialize(root))
	t.Log(droot.Val)
	t.Log(droot.Left.Val)
	t.Log(droot.Right.Val)
	t.Log(droot.Right.Left.Val)
	t.Log(droot.Right.Right.Val)

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
