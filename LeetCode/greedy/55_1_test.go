package greedy

import (
	"testing"
)

func canJump(nums []int) bool {
	var maxJump int

	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}

		maxJump = max(maxJump, i+nums[i])
		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestNums(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	t.Log(canJump(nums))
}
