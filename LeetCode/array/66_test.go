package array

import (
	"testing"
)

func plusOne(digits []int) []int {
	for tail := len(digits) - 1; tail >= 0; tail-- {
		digits[tail] = digits[tail] + 1
		if digits[tail] == 10 {
			digits[tail] = 0
			if tail == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		} else {
			return digits
		}
	}
	return digits

}

func TestPlusOne(t *testing.T) {
	digits := []int{1, 2, 3}
	t.Log(plusOne(digits))
}
