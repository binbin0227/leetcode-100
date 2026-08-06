package main

func searchRange(nums []int, target int) []int {
	findStart := func() int {
		start := -1
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2
			if target > nums[mid] {
				l = mid + 1
			} else if target < nums[mid] {
				r = mid - 1
			} else {
				start = mid
				r = mid - 1
			}
		}
		return start
	}

	findEnd := func() int {
		end := -1
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2
			if target > nums[mid] {
				l = mid + 1
			} else if target < nums[mid] {
				r = mid - 1
			} else {
				end = mid
				l = mid + 1
			}
		}
		return end
	}

	return []int{findStart(), findEnd()}
}
