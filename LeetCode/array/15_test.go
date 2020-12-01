package array

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	length := len(nums)
	for index, value := range nums {
		if nums[index] > 0 {
			return res
		}
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		l := index + 1
		r := length - 1
		for l < r {
			sum := value + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{nums[index], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			case sum > 0:
				r -= 1
			case sum < 0:
				l += 1
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	num := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(num))
}
