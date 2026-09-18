package main

func nextPermutation(nums []int) {
	// 1`3`-5-`4`2 ->14(532) -> 14235

	n := len(nums)
	if n == 1 {
		return
	}

	reverse := func(left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	i := 0
outer:
	for i = n - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			for j := n - 1; j >= 1; j-- {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
					break outer
				}
			}
		}
	}
	reverse(i, n-1)
}