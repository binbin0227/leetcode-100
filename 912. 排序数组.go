package main

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
    return nums
}

func sort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2
	sort(nums, l, mid)
	sort(nums, mid+1, r)

	temp := make([]int, r-l+1)
	curr := 0
	p1, p2 := l, mid+1
	for p1 <= mid && p2 <= r {
		if nums[p1] < nums[p2] {
			temp[curr] = nums[p1]
			p1++
		} else {
			temp[curr] = nums[p2]
			p2++
		}
		curr++
	}
	for p1 <= mid {
		temp[curr] = nums[p1]
		p1++
		curr++
	}
	for p2 <= r {
		temp[curr] = nums[p2]
		p2++
		curr++
	}

	copy(nums[l: r+1], temp) // 左闭右开

    return
}