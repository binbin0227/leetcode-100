package main

func lengthOfLIS(nums []int) int {
	// F(i) 表示以第 i 个数字作为结尾的递增子序列
	// F(i) = max(F(j) + 1, F(i)) (nums[i] > nums[j])
	dp := make([]int, len(nums))

	for i := range len(nums) {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := range i {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	var res int
	for _, count := range dp {
		res = max(res, count)
	}
	return res
}
