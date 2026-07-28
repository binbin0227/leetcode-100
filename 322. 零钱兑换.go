package main

func coinChange(coins []int, amount int) int {
	// f(i) = f(i-j) + 1
	dp := make([]int, amount+1)

	// 初始设为无穷大
	for i := range amount + 1 {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i { // 防越界
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
