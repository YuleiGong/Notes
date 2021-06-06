package dfs_dfs

import "testing"

func numIslands(grid [][]byte) int {
	var num int
	var width = len(grid[0])
	var hight = len(grid)
	var dfs func(h, w int)

	dfs = func(h, w int) {
		if w >= 0 && h >= 0 && w < width && h < hight && grid[h][w] == '1' {
			grid[h][w] = '0'
			dfs(h, w+1)
			dfs(h, w-1)
			dfs(h+1, w)
			dfs(h-1, w)
		}
	}

	for w := 0; w < width; w++ {
		for h := 0; h < hight; h++ {
			if grid[h][w] == '1' {
				num++
				dfs(h, w)
			}
		}
	}

	return num
}

func TestNumIsLands(t *testing.T) {
	var grid = [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	t.Logf("%v", numIslands(grid))
}
