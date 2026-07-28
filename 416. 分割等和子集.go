package main

func canPartition(nums []int) bool {
	// F(i) 表示：是否能从数组中挑出一些数，刚好凑出和为 i
	// F(i) = F(i) || F(i - nums[j])
	// 如果 sum 为奇数，直接返回 false
	// 零件不能重复使用，所以双层循环倒着来,先拿零件，再拿数
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)

	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}
