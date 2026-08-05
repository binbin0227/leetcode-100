package main

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, n)
	for i := range n {
		path[i] = make([]byte, n)
		for j := range n {
			path[i][j] = '.'
		}
	}

	check := func(x, y int) bool {
		// 同一行
		for j := range y {
			if path[x][j] == 'Q' {
				return false
			}
		}
		// 同一列
		for i := range x {
			if path[i][y] == 'Q' {
				return false
			}
		}
		// 对角线
		for i := range x {
			if y-x+i >= 0 && path[i][y-x+i] == 'Q' {
				return false
			}
			if y+x-i < n && path[i][y+x-i] == 'Q' {
				return false
			}
		}

		return true
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, 0)
			for i := range len(path) {
				temp = append(temp, string(path[i]))
			}
			res = append(res, temp)
			return
		}

		for col := range n {
			if check(row, col) {
				path[row][col] = 'Q'
				backtrack(row + 1)
				path[row][col] = '.'
			}
		}
	}

	backtrack(0)
	return res
}
