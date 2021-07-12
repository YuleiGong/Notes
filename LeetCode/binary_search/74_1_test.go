package main

import (
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix[0])-1

	for _, nums := range matrix {
		if nums[right] >= target { //执行二分
			for left <= right {
				mid := left + (right-left)/2
				if nums[mid] == target {
					return true
				}
				if nums[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			return false
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T) {
	/*
		matrix := [][]int{
			[]int{1, 3, 5, 7},
			[]int{10, 11, 16, 20},
			[]int{23, 30, 34, 60},
		}*/
	matrix := [][]int{
		[]int{1},
		[]int{3},
	}

	t.Log(searchMatrix(matrix, 0))
}
