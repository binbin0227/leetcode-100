package main

func findDuplicates(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)

	idx := 0
	for idx < n {
		curr := nums[idx]

		// 已经找到了的、记录过的且不在位置上的数标记成0
		if curr == 0 {
			idx++
			continue
		}

		// curr 如果已经就位就跳过
		if curr == idx+1 {
			idx++
			continue
		}

		// curr 要去的位置已经有人了，说明重复了
		if nums[curr-1] == curr {
			res = append(res, curr)
			nums[idx] = 0 // 标记
			idx++
			continue
		}

		nums[idx], nums[curr-1] = nums[curr-1], nums[idx]
	}

	return res
}
