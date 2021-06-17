package dfs_dfs

import "testing"

func numIslands(grid [][]byte) int {
	xCount := len(grid)
	yCount := len(grid[0])
	var count int
	var dfs func(x, y int)

	dfs = func(x, y int) {
		if x >= 0 && y >= 0 && x < xCount && y < yCount && grid[x][y] == '1' {
			grid[x][y] = '0'
			dfs(x+1, y)
			dfs(x-1, y)
			dfs(x, y+1)
			dfs(x, y-1)
		}
	}

	for x := 0; x < xCount; x++ {
		for y := 0; y < yCount; y++ {
			if grid[x][y] == '1' {
				count++
				dfs(x, y)
			}
		}
	}
	return count
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
