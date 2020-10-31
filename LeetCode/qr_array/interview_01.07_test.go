package qr_array

import (
	"testing"
)

// https://leetcode-cn.com/problems/rotate-matrix-lcci/solution/golang-shuang-yi-bai-guan-fang-ti-jie-fang-fa-3-by/
func rotate(matrix [][]int) {
	//先水平翻转
	n := len(matrix)
	for y := 0; y < n/2; y++ {
		matrix[y], matrix[n-1-y] = matrix[n-1-y], matrix[y]
	}

	//对角线翻转
	for y := 0; y < n; y++ {
		for x := 0; x < y; x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}
}

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	t.Log(matrix)
}
