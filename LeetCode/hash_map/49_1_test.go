package hash_map

import (
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		hash[string]
	}

}

func TestGroupAnarams(t *testing.T) {
	strs := []stirng{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(strs))
}
