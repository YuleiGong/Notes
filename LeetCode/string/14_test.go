package string

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	prefix := strs[0] //使用第一个元素最为最长前缀
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix

}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"hello", "hello world"}
	t.Log(longestCommonPrefix(strs))
}
