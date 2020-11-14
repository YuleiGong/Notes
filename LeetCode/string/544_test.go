package string

import (
	"testing"
)

func reverseWords(s string) string {
	byteS := []byte(s)
	length := len(s)

	start := 0
	for i, v := range s {
		if v == ' ' || i == length-1 {
			end := i - 1
			if i == len(s)-1 { //最后一个单词
				end = i
			}
			for start < end {
				byteS[start], byteS[end] = byteS[end], byteS[start] //原地交换
				start++
				end--
			}
			start = i + 1
		}
	}
	return string(byteS)
}

func TestReverseWords(t *testing.T) {
	//s := "Let's take LeetCode contest"
	s := "Let's "
	t.Log(reverseWords(s))
}
