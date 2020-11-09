package string

import (
	"testing"
)

func reverseString(s []byte) {
	length := len(s)
	start := 0
	end := length - 1
	if length == 1 {
		return
	}

	for {
		//偶数
		if length%2 == 0 && start == length/2 {
			break
		}
		s[end], s[start] = s[start], s[end] //交换
		end--
		start++
		//奇数
		if length%2 != 0 && end == start {
			break
		}

	}
}

func TestReverseString(t *testing.T) {
	str := "h"
	reverseString([]byte(str))
}
