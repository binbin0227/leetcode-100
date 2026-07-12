package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	var dfs func(idx, i, j int) bool
	dfs = func(idx, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = '#'
		res := dfs(idx+1, i+1, j) || dfs(idx+1, i-1, j) || dfs(idx+1, i, j+1) || dfs(idx+1, i, j-1)
		board[i][j] = temp
		return res
	}
	for i := range m {
		for j := range n {
			if board[i][j] == word[0] {
				if dfs(0, i, j) {
					return true
				}
			}
		}
	}
	return false
}
