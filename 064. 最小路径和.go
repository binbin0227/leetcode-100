package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}
