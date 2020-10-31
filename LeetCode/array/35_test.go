package array

import "testing"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}

		if i+1 == len(nums) {
			return i + 1
		}
		if nums[i] < target && nums[i+1] > target {
			return i + 1
		}

	}
	return len(nums)
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	t.Log(searchInsert(nums, 0))
}
