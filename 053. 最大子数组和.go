package main

import "math"

func maxSubArray(nums []int) int {
	// 只要遇到负数，就把 sum 清零
	// 先维护结果再清零，否则要是 sum 全为负数 res 得不到维护就会返回无限小
	res := math.MinInt
	sum := 0
	for i := range len(nums) {
		sum += nums[i]
		res = max(res, sum)
		if sum < 0 {
			sum = 0
			continue
		}
	}
	return res
}
