package main

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	i := 0

	for i <= p2 {
		if nums[i] == 0 && p0 <= i {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		} else if nums[i] == 2 && i <= p2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		} else {
			i++
		}
	}
}
