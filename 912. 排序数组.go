package main

func sortArray(nums []int) []int {
    // 把排序好的两边放到 temp 里面，再复制回原数组

	temp := make([]int, len(nums))

	var sort func(left, right int)
	sort = func(left, right int) {
		if left >= right {
			return
		}

		mid := left + (right-left)/2
		sort(left, mid)
		sort(mid+1, right)

		pLeft := left
		pRight := mid + 1
		pTemp := left

		for pLeft <= mid && pRight <= right {
			if nums[pLeft] < nums[pRight] {
				temp[pTemp] = nums[pLeft]
				pLeft++
			} else {
				temp[pTemp] = nums[pRight]
				pRight++
			}
			pTemp++
		}
		for pLeft <= mid {
			temp[pTemp] = nums[pLeft]
			pLeft++
			pTemp++
		}
		for pRight <= right {
			temp[pTemp] = nums[pRight]
			pRight++
			pTemp++
		}

		for p := left; p <= right; p++ {
			nums[p] = temp[p]
		}
	}

	sort(0, len(nums)-1)
	return nums
}