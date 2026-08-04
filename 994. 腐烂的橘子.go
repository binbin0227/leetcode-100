package main

func orangesRotting(grid [][]int) int {
	minutes := 0
	freshCount := 0
	m, n := len(grid), len(grid[0])
	q := make([][]int, 0)

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				freshCount++
			}
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}

	var infect func(i, j int)
	infect = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			grid[i][j] = 2
			freshCount--
			q = append(q, []int{i, j})
		}
	}
	var spread func(i, j int)
	spread = func(i, j int) {
		infect(i+1, j)
		infect(i-1, j)
		infect(i, j+1)
		infect(i, j-1)
	}

	for freshCount != 0 && len(q) != 0 {
		minutes += 1
		length := len(q)
		for idx := range length {
			badLoc := q[idx]
			spread(badLoc[0], badLoc[1])
		}
		q = q[length:]
	}

	if freshCount == 0 {
		return minutes
	}
	return -1
}
