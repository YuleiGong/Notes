package array

import (
	"testing"
)

func moveZeroes(nums []int) {
	lNonZ := 0 //记录最近一次等于0的下标
	for i := range nums {
		if nums[i] != 0 {
			if i != lNonZ { //连续数字
				nums[lNonZ] = nums[i]
				nums[i] = 0
			}
			lNonZ++
		}
	}
}

func TestMoveZeros(t *testing.T) {
	nums := []int{1, 1}
	moveZeroes(nums)
}
