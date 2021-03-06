package recursion

import "testing"

func permute(nums []int) [][]int {
	var result [][]int
	visited := make(map[int]bool)
	var path []int
	var helper func()

	helper = func() {
		if len(nums) == len(path) {
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
			helper()                  //每一层遍历
			path = path[:len(path)-1] //撤销最后一个无效值
			visited[n] = false
		}
	}

	helper()

	return result
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	permute()
}
