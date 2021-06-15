package recursion

import "testing"

func permute(nums []int) [][]int {
	var result [][]int
	var helper func()
	var path []int
	visited := make(map[int]bool)

	helper = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)

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
