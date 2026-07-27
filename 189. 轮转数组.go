package main

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	temp := nums[n-k:]
	temp = append(temp, nums[:n-k]...)
	copy(nums, temp)
}
