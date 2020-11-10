package array

import "testing"

//双指针
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
		if sum == target {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {
	nums := []int{5, 25, 75}
	target := 100
	t.Log(twoSum(nums, target))
}
