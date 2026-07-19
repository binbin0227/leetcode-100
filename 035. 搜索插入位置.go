package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
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
	return left
}

func searchInsert2(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}
func binarySearch(nums []int, target, left, right int) int {
	if left > right {
		return left
	}
	mid := left + (right-left)/2
	if target < nums[mid] {
		return binarySearch(nums, target, left, mid-1)
	} else if target > nums[mid] {
		return binarySearch(nums, target, mid+1, right)
	} else {
		return mid
	}
}
