package main

func moveZeroes(nums []int) {
	r := 0
	for l := range len(nums) {
		if nums[l] != 0 {
			continue
		}
		if r <= l {
			r = l + 1
		}
		for r < len(nums) && nums[r] == 0 {
			r++
		}
		if r == len(nums) {
			return
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
}
