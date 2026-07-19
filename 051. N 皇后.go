package main

import "strings"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	var check func(path []string, row, col int) bool
	check = func(path []string, row, col int) bool {
		for i := range row {
			if path[i][col] == 'Q' {
				return false
			}
		}
		i, j, k := row-1, col-1, col+1
		for i >= 0 {
			if j >= 0 && path[i][j] == 'Q' {
				return false
			}
			if k < n && path[i][k] == 'Q' {
				return false
			}
			i--
			j--
			k++
		}
		return true
	}
	var backtrack func(path []string, n, row int)
	backtrack = func(path []string, n, row int) {
		if row == n {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for col := range n {
			if check(path, row, col) {
				path = append(path, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
				backtrack(path, n, row+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(nil, n, 0)
	return res
}
