package recursion

import (
	"fmt"
	"testing"
)

func combine(n int, k int) [][]int {
	var result [][]int
	var helper func(start int, path []int)

	helper = func(start int, path []int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			helper(i+1, path)
			fmt.Println(path)
			path = path[:len(path)-1]
		}
	}

	helper(1, []int{})

	return result
}

func TestCombine(t *testing.T) {
	t.Log(combine(4, 2))
}
