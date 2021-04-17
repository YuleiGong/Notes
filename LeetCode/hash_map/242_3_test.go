package main

import "testing"

func isAnagram(s string, t string) bool {
	hash := make(map[byte]int64)
	for _, c := range []byte(s) {
		hash[c]++
	}

	for _, c := range []byte(t) {
		hash[c]--
		if hash[c] < 0 {
			return false
		}
	}
	return true

}

func TestIsAnagram(t *testing.T) {
	s := "rat"
	t1 := "car"
	t.Log(isAnagram(s, t1))
}
