package greedy

import "testing"

func jump(nums []int) int {
	var maxPos int
	var count int
	var end int

	for i := 0; i < len(nums)-1; i++ {
		maxPos := max(maxPos, nums[i]+i) //任意一点所能跳跃的最大位置

		if i == end { //跳一次 尽可能的走最多在跳
			end = maxPos
			count++
		}
	}

	return count

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	t.Log(jump(nums))
}
