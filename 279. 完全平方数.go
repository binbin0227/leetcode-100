package main

func numSquares(n int) int {
	// F(i) = min(F(i-j^2)) + 1

	// 先假设最坏情况，n 要经过 n 个 1 来拼接
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = i
	}

	for i := 2; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}
