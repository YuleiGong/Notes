package array

import "testing"

//头尾指针，逐步逼近
func maxArea(height []int) int {
	max := 0
	var area int
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
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
}
