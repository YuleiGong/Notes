package array

import (
	"testing"
)

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	left := 0
	for j := 0; j < len(nums); j++ {
		if left == (sum - left - nums[j]) {
			return j
		}
		left += nums[j]
	}
	return -1
}

func TestPivotIndex(t *testing.T) {
	//nums := []int{1, 7, 3, 6, 5, 6}
	//nums := []int{1, 2, 3}
	nums := []int{-1, -1, -1, 0, 1, 1}
	t.Log(pivotIndex(nums))
}
