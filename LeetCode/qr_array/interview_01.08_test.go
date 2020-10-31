package qr_array

import (
	"testing"
)

func setZeroes(matrix [][]int) {
	colLength := len(matrix[0])
	rowLength := len(matrix)
	//使用map  记录需要置0的行列
	colClear := make(map[int]bool)
	rowClear := make(map[int]bool)

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			if matrix[i][j] == 0 {
				if _, ok := colClear[j]; !ok {
					colClear[j] = true
				}
				if _, ok := rowClear[i]; !ok {
					rowClear[i] = true
				}
			}
		}
	}

	for col, _ := range colClear {
		for i := 0; i < rowLength; i++ {
			matrix[i][col] = 0
		}
	}

	for row, _ := range rowClear {
		for i := 0; i < colLength; i++ {
			matrix[row][i] = 0
		}
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{
		{0, 1},
	}
	/*
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
	*/
	setZeroes(matrix)
	t.Log(matrix)
}
