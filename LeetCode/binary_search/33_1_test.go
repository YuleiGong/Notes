package main

import "testing"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] { //
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //mid 在右侧
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}

		}
	}

	return -1

}
func TestSearch(t *testing.T) {
	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

}
