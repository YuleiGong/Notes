package binary_search

import "testing"

func mySqrt(x int) int {
	var start int = 0
	var end int = x
	var result = 0

	for start <= end {
		mid := start + (end-start)/2
		if mid*mid <= x {
			result = mid //只会去整数部分
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result

}

func TestMySqrt(t *testing.T) {
	t.Log(mySqrt(32343434))
}
