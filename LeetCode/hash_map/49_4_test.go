package hash_map

import (
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	hashMap := make(map[string][]string)

	for _, v := range strs {
		bs := []byte(v)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		hashMap[string(bs)] = append(hashMap[string(bs)], v)
	}
	for _, v := range hashMap {
		result = append(result, v)
	}

	return result

}
func TestGroupAnarams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(strs))
}
