package main

func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}