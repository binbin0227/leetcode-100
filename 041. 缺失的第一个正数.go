package main

func firstMissingPositive(nums []int) int {
	// 用自己当哈希表，如 1 放到索引 0 的位置
	n := len(nums)
	for i := range n {
		// 交换条件：0 < num < len(nums) 且和要交换的数不等
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := range n {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
