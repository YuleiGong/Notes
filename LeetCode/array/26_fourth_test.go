package array

import "testing"

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1

}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(nums))

}
