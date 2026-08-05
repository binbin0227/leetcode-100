package main

func exist(board [][]byte, word string) bool {
	res := false
	m, n := len(board), len(board[0])

	var backtrack func(startIdx int, i, j int) bool
	backtrack = func(startIdx, i, j int) bool {
		if startIdx == len(word) {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		if board[i][j] != word[startIdx] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'
		flag := backtrack(startIdx+1, i+1, j) ||
			backtrack(startIdx+1, i-1, j) ||
			backtrack(startIdx+1, i, j+1) ||
			backtrack(startIdx+1, i, j-1)
		board[i][j] = temp

		return flag
	}

outer:
	for i := range m {
		for j := range n {
			if board[i][j] == word[0] {
				res = backtrack(0, i, j)
			}
			if res == true {
				break outer
			}
		}
	}

	return res
}
