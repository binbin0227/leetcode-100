package main

func searchRange(nums []int, target int) []int {
	leftIdx := searchLeft(nums, target)
	rightIdx := searchRight(nums, target)
	if leftIdx > rightIdx || leftIdx >= len(nums) || rightIdx < 0 {
		return []int{-1, -1}
	}
	return []int{leftIdx, rightIdx}
}
func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target <= nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
func searchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
