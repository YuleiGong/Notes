package qr_array

import (
	"testing"
)

//遍历打印出二维数组对角线遍历的值
//下对角线 x增加 y减小
type prevXY struct {
	row int
	col int
}

func findDiagonalOrder(matrix [][]int) []int {
	colLen := len(matrix[0])
	rowLen := len(matrix)
	isDown := true
	isUp := false
	result := make([]int, 0)
	prev := &prevXY{0, 0} //起点
	result = append(result, matrix[prev.row][prev.col])
	prev.col++

	for {
		if isDown {
			prev.row++
			isUp = true
			isDown = false
			for {
				result = append(result, matrix[prev.row][prev.col])
				prev.row++
				prev.col--
				if prev.col < 0 || prev.row > rowLen-1 {
					prev.row--
					prev.col++
					break
				}
			}
		}
		if isUp {
			prev.row++
			isDown = true
			isUp = false
			for {
				result = append(result, matrix[prev.row][prev.col])
				prev.row--
				prev.col++
				if prev.row < 0 || prev.col > colLen-1 {
					prev.row++
					prev.col--
					break
				}
			}
		}
	}
}

func TestFindDiagonalOrder(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Log(findDiagonalOrder(array))
}
