package main

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])

	var clear func(grid [][]byte, i, j int)
	clear = func(grid [][]byte, i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] != '1' {
			return
		}

		grid[i][j] = '#'
		clear(grid, i+1, j)
		clear(grid, i-1, j)
		clear(grid, i, j+1)
		clear(grid, i, j-1)
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				res += 1
				clear(grid, i, j)
			}
		}
	}

	return res
}
