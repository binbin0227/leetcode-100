package main

func maxArea(height []int) int {
	// 双指针初始指向两端，谁的高度矮，谁就往里缩,每缩一次都要维护最大水量
	var res int
	l := 0
	r := len(height) - 1

	for l < r {
		if height[l] <= height[r] {
			water := (r - l) * height[l]
			if water > res {
				res = water
			}
			l++
		} else {
			water := (r - l) * height[r]
			if water > res {
				res = water
			}
			r--
		}
	}
	return res
}
