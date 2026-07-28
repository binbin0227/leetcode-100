package main

func rob(nums []int) int {
	// F(i) = max(F(i-2)+nums[i], F(i-1))
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	var res int
	f2, f1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		res = max(f1, f2+nums[i])
		f2, f1 = f1, res
	}

	return res
}
