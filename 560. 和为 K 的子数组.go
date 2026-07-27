package main

func subarraySum(nums []int, k int) int {
	// 滑动窗口做法只适用于数组内全为正数
	// 前缀表: 前缀和 = 出现过的次数
	// 当前位置的前缀和减掉 k 得到的前缀和出现的次数就是子串的出现次数

	res := 0
	preSumMap := make(map[int]int, 0)
	preSum := 0
	preSumMap[0] = 1 // 前缀和为0的情况出现过一次

	for _, num := range nums {
		preSum += num

		if times, ok := preSumMap[preSum-k]; ok {
			res += times
		}

		preSumMap[preSum]++

	}

	return res
}
