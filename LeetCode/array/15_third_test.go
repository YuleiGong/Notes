package array

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for base := 0; base < len(nums); base++ {
		if nums[base] > 0 {
			return result
		}
		if base > 0 && nums[base] == nums[base-1] {
			continue
		}
		head := base + 1
		tail := len(nums) - 1
		for head < tail {
			target := nums[base] + nums[head] + nums[tail]
			switch {
			case target == 0:
				result = append(result, []int{nums[base], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] {
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			case target < 0:
				head++
			case target > 0:
				tail--
			}

		}
	}

	return result
}

func TestThreeSum(t *testing.T) {
	num := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(num))
}
