package array

import (
	"testing"
)

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	prev := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 && prev != 1 {
			if count >= maxCount {
				maxCount = count
			}
			count = 1
		}
		if nums[i] == 1 && prev == 1 {
			count++
		}
		prev = nums[i]
	}

	if count >= maxCount {
		maxCount = count
	}

	return maxCount

}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{0, 0, 0}
	t.Log(findMaxConsecutiveOnes(nums))
}
