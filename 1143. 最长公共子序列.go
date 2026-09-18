package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 表示字符串 1 的前 i 个字符和字符串 2 的前 j 个字符的最长公共子序列长度
	// dp[i][j] = dp[i-1][j-1] + 1           if text1[i] == text2[j]
	//            max(dp[i-1][j],dp[i][j-1])

	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
