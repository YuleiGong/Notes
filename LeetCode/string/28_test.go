package string

import "testing"

func strStr(haystack string, needle string) int {
	//注意处理特殊条件
	if haystack == "" && needle == "" {
		return 0
	}
	if haystack == "" && needle != "" {
		return -1
	}

	length := len(haystack)
	for i := 0; i < length; i++ {
		end := i + len(needle)
		if end > length {
			return -1
		}
		if haystack[i:end] == needle {
			return i
		}
	}
	return -1

}

func TestStrStr(t *testing.T) {
	haystack := "mississippo"
	needle := "issi"
	t.Log(strStr(haystack, needle))
}
