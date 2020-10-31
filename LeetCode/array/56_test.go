package array

import (
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else if intervals[i][1] > result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result

}

func TestMerge(t *testing.T) {
	//nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//nums := [][]int{{1, 4}, {4, 5}}
	//nums := [][]int{{1, 4}, {4, 5}, {2, 3}}
	nums := [][]int{{1, 4}, {0, 1}}
	t.Log(merge(nums))
}
