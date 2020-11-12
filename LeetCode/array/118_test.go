package array

import "testing"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}

	for i := 3; i <= numRows; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for j := 0; j < len(res[i-2])-1; j++ {
			tmp = append(tmp, res[i-2][j]+res[i-2][j+1])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}

	return res
}

func TestGenerate(t *testing.T) {
	t.Log(generate(0))
}
