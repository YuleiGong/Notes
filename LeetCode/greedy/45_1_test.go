package greedy

import "testing"

func jump(nums []int) int {
	var count int
	var end int
	var maxPos int

	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)

		if i == end { //end 可以到达的最远距离
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
