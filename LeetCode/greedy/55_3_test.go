package greedy

import (
	"testing"
)

func canJump(nums []int) bool {
	var maxJump int

	for i := 0; i < len(nums); i++ {
		if i > maxJump { //当前位置小于所能跳跃的最大位置
			return false
		}

		maxJump = max(maxJump, nums[i]+i) //所能跳到的最远距离
		if maxJump >= len(nums)-1 {       //len(nums) - 1 剩余最远距离
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
	nums := []int{3, 2, 1, 0, 4}
	t.Log(canJump(nums))
}
