package main

import "sort"

func threeSum(nums []int) [][]int {
	// 先给数组排序，然后固定一个数，它的下一个数为左指针，最后一个数为右指针，然后求和。
	// 和小于0则右移左指针，大于0则左移右指针
	// 找到符合条件之后要继续找
	// 记得去重
	var res [][]int
	sort.Ints(nums)

	for i := range len(nums) - 1 {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue // 去重
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l != i+1 && nums[l] == nums[l-1] {
				l++
				continue // 去重
			}
			if r != len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue // 去重
			}

			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return res
}
