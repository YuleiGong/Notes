package array

import "testing"

func getRow(rowIndex int) []int {
	nums := []int{1}

	for i := 1; i <= rowIndex; i++ {
		nums = append(nums, 1)
		for j := i - 1; j > 0; j-- {
			nums[j] += nums[j-1]
		}
	}
	return nums
}

func TestGetRow(t *testing.T) {
	t.Log(getRow(3))

}
