package stack

import (
	"testing"
)

func isValid(s string) bool {
	hash := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]byte, 0)
	for _, c := range []byte(s) {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) > 0 && c == hash[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false //eg: a := '}'
		}
	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	a := "}"
	t.Log(isValid(a))

}
