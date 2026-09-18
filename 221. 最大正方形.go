package main

func maximalSquare(matrix [][]byte) int {
	//dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	maxLen := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen * maxLen
}
