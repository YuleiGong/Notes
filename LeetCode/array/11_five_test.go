package array

import "testing"

func maxArea(height []int) int {
	var area int
	var max int

	tail := len(height) - 1
	for head := 0; head < tail; {
		if height[head] < height[tail] {
			area = (tail - head) * height[head]
			head++
		} else {
			area = (tail - head) * height[tail]
			tail--
		}
		if max < area {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	height := []int{4, 3, 2, 1, 4}
	t.Log(maxArea(height))
}
