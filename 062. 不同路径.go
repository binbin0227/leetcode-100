package main

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i-1][j]+dp[i][j-1]
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
