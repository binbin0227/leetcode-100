package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if text1[i] == text2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
					continue
				}
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 0
					continue
				}
				if i == 0 {
					dp[i][j] = dp[i][j-1]
					continue
				}
				if j == 0 {
					dp[i][j] = dp[i-1][j]
					continue
				}
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
