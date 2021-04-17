package array

import "testing"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, val := range nums {
		if p, ok := hash[target-val]; ok {
			return []int{p, i}
		}
		hash[val] = i
	}

	return []int{}

}

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	t.Log(twoSum(nums, target))

}
