package array

import (
	"sort"
	"testing"
)

//先做一次排序
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func TestArrayPairSum(t *testing.T) {
	nums := []int{6, 2, 6, 5, 1, 2}
	t.Log(arrayPairSum(nums))
}
