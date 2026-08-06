package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] { // 左边有序
			if target >= nums[left] && target < nums[mid] { // 判断是不是在有序区间里
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
