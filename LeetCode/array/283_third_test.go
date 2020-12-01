package array

import "testing"

func moveZeroes(nums []int) {
	low := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if low != i {
				nums[low] = nums[i]
				nums[i] = 0
			}
			low++
		}
	}
}

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
