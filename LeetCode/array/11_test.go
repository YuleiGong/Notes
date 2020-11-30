package array

import "testing"

func maxArea(height []int) int {
	var area int
	var max int
	tail := len(height) - 1 //尾指针 头尾指针，分别向内收敛
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
	height := []int{1, 2, 1}
	t.Log(maxArea(height))
}
