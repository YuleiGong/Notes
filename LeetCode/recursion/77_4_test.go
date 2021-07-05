package recursion

import (
	"testing"
)

func combine(n int, k int) [][]int {
	var result [][]int
	var helper func(start int, val []int)

	helper = func(start int, val []int) {
		if len(val) == k {
			tmp := make([]int, k)
			copy(tmp, val)
			result = append(result, tmp)
		}

		for i := start; i <= n; i++ {
			val = append(val, i)
			helper(i+1, val)
			val = val[:len(val)-1]
		}
	}

	helper(1, []int{})

	return result
}

func TestCombine(t *testing.T) {
	t.Log(combine(4, 2))
}
