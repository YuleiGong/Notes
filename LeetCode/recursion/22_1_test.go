package recursion

import (
	"testing"
)

func generateParenthesis(n int) []string {
	var result []string
	var helper func(left, right int, path string)

	helper = func(left, right int, path string) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		//校验括号的合法性
		if left < n { //一直添加左括号，只要别用完就行
			helper(left+1, right, path+"(")
		}

		if left > right { //左括号 一定会比右括号多
			helper(left, right+1, path+")")
		}
	}
	helper(0, 0, "")
	return result
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}
