package recursion

import (
	"testing"
)

func generateParenthesis(n int) []string {
	var result []string
	var path string
	var helper func(left, right int, path string)

}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(2))
}
