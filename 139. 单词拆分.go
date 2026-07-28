package main

func wordBreak(s string, wordDict []string) bool {
	// F(i) = F(j) && s[j+1:] -> wordDict
	dp := make([]bool, len(s)+1)

	dp[0] = true
	for i := 1; i <= len(s); i++ { // 这里的 1 表示第一位，也就是索引 0
		for _, word := range wordDict {
			// 零件小于整体 && F(j) true && s[j+1:] -> wordDict
			if len(word) <= i && dp[i-len(word)] && s[i-len(word):i] == word {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
