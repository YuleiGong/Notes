package recursion

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
	var s string
	var helper func(*TreeNode)

	helper = func(*TreeNode) {

	}

}

// Deserializes your encoded data to tree. 反序列化
func (this *Codec) deserialize(data string) *TreeNode {

}

func TestEncodeOrDecode() {
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
