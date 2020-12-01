package array

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := make([][]int, 0)

	for base := 0; base < length; base++ {
		if nums[base] > 0 {
			return res
		}
		//去重
		if base > 0 && nums[base] == nums[base-1] {
			continue
		}
		head := base + 1
		tail := length - 1
		for head < tail {
			sum := nums[base] + nums[head] + nums[tail]
			switch {
			case sum == 0:
				res = append(res, []int{nums[base], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] {
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			case sum < 0:
				head++
			case sum > 0:
				tail--
			}
		}
	}

	return res
}

func TestThreeSum(t *testing.T) {
	num := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(num))
}
