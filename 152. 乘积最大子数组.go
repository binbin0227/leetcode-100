package main

import "math"

func maxProduct(nums []int) int {
	currMax := nums[0]
	currMin := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		currMax, currMin = max(nums[i], nums[i]*currMax, nums[i]*currMin), min(nums[i], nums[i]*currMax, nums[i]*currMin)
		res = max(res, currMax)
	}

	return res
}

func maxProduct2(nums []int) int {
	// F(i) = max(F(j) * nums[j+1]*...*nums[i], F(i), nums[j+1]*...*nums[i])
	dp := make([]int, len(nums))

	for i := range len(nums) {
		dp[i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		for j := range i {
			temp := 1
			for k := j + 1; k <= i; k++ {
				temp *= nums[k]
			}
			dp[i] = max(dp[i], dp[j]*temp, temp)
		}
	}

	res := math.MinInt
	for _, p := range dp {
		res = max(res, p)
	}
	return res
}
