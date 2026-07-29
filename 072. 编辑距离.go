package main

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}

	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if word1[i] == word2[j] {
				if i == 0 && j == 0 {
					dp[i][j] = 0
					continue
				}
				if i == 0 {
					dp[i][j] = j
					continue
				}
				if j == 0 {
					dp[i][j] = i
					continue
				}
				dp[i][j] = dp[i-1][j-1]
			} else {
				if i == 0 && j == 0 {
					dp[i][j] = 1
					continue
				}
				if i == 0 {
					dp[i][j] = dp[i][j-1] + 1
					continue
				}
				if j == 0 {
					dp[i][j] = dp[i-1][j] + 1
					continue
				}
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m-1][n-1]
}
