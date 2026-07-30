package main

func nextPermutation(nums []int) {
	// 13542 ->14532 -> 14235
	n := len(nums)
	var i int

	for i = n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		for j := n - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}

	l, r := i+1, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
