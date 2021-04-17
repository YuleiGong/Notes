package array

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) //小到大
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
			sum := nums[base] + nums[head] + nums[tail]
			switch {
			case sum == 0:
				result = append(result, []int{nums[base], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] { //左右移动，如果有重复的，需要多次移动
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			case sum > 0:
				tail--
			case sum < 0:
				head++
			}
		}
	}

	return result

}

func TestThreeSum(t *testing.T) {
	num := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(num))
}
