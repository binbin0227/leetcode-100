package main

func firstMissingPositive(nums []int) int {
	// 用自己当哈希表，如 1 放到索引 0 的位置
	n := len(nums)
	for i := range n {
		// 这里要用 for，因为把 nums[i] 归位之后，新来的 nums[i] 可能还不在正确位置
		// 而且要防止 nums[i] 和它打算换的数相等的死循环情况
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
