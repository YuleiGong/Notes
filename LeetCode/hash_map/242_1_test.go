package hash_map

import (
	"testing"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[rune]int64)
	for _, c := range s {
		hash[c]++
	}

	for _, c := range t {
		hash[c]--
		if hash[c] < 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	t1 := "nagaram"
	t.Log(isAnagram(s, t1))
}
