package hash_map

import (
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	result := make([][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) //排序
		hash[string(s)] = append(hash[string(s)], str)
	}

	for _, v := range hash {
		result = append(result, v)
	}

	return result
}

func TestGroupAnarams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(strs))
}
