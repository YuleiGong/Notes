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

		if left < n {
			helper(left+1, right, path+"(")
		}

		if right < left {
			helper(left, right+1, path+")")
		}
	}

	helper(0, 0, "")

	return result
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(2))
}
