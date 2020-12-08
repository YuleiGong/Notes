package array

import "testing"

func moveZeroes(nums []int) {
	zeroIndex := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != zeroIndex {
				nums[zeroIndex] = nums[i]
				nums[i] = 0
			}
			zeroIndex++
		}
	}
}

func TestMoveZero(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
