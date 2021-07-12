package main

import (
	"testing"
)

func isPerfectSquare(num int) bool {
	var start, end int = 1, num

	for start <= end { //结束条件
		mid := start + (end-start)/2
		switch {
		case mid*mid == num:
			return true
		case mid*mid < num:
			start = mid + 1
		case mid*mid > num:
			end = mid - 1
		}

	}

	return false
}

func TestIsPerfectSquare(t *testing.T) {
	t.Logf("%v", isPerfectSquare(13))
}
