package array

import "testing"

func moveZeroes(nums []int) {
	low := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[low] = nums[i]
			if i != low {
				nums[i] = 0
			}
			low++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{1, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
