package array

import "testing"

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[0 : k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	end := len(nums) - 1
	for start := 0; start < len(nums)/2; start++ {
		nums[start], nums[end] = nums[end], nums[start]
		end--
	}
}

func TestRotate(t *testing.T) {
	nums := []int{-1}
	rotate(nums, 2)
	t.Log(nums)
}
