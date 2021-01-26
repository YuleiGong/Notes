package stack

import (
	"testing"
)

func isValid(s string) bool {
	hash := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) > 0 && hash[stack[len(stack)-1]] == c {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	a := "({)"
	t.Log(isValid(a))

}
