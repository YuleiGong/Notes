package binary_search

import "testing"

func mySqrt(x int) int {
	var start int = 0
	var end int = x
	var result int

	for start <= end {
		mid := start + (end-start)/2
		if mid*mid <= x {
			result = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result

}

func TestMySqrt(t *testing.T) {
	t.Log(mySqrt(9))
}
