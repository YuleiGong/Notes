package recursion

import "testing"

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	visited := make(map[int]bool)
	var helper func()

	helper = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			helper()
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	helper()

	return result
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	t.Log(permute(nums))
}
