package main

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	mayBeMax := make([]int, 0)

	// 直接一个循环搞定，不需要显式维护 l 和 r
	for i := range len(nums) {

		// 只要切片不为空，且当前新人 nums[i] >= 尾部对应的值，就无情踢掉尾部
		for len(mayBeMax) > 0 && nums[i] >= nums[mayBeMax[len(mayBeMax)-1]] {
			mayBeMax = mayBeMax[:len(mayBeMax)-1]
		}

		mayBeMax = append(mayBeMax, i)

		// 清理过期数据
		if mayBeMax[0] < i-k+1 {
			mayBeMax = mayBeMax[1:]
		}

		if i >= k-1 {
			res = append(res, nums[mayBeMax[0]])
		}
	}

	return res
}
