package hash_map

import (
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	result := make([][]string, 0)

	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] }) //排序 原地排序
		hash[string(s)] = append(hash[string(s)], v)
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
