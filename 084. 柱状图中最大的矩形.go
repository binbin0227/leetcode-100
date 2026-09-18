package main

func largestRectangleArea(heights []int) int {
	// 单调递增栈；当遇到更矮柱子时，弹出的柱子终于确定了左右边界，此时计算它能形成的最大矩形。

	res := 0
	stack := make([]int, 0)
	newHeights := make([]int, len(heights)+2)
	copy(newHeights[1:], heights) // 在数组的前面和后面插入 0

	for i := range len(newHeights) {
		for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
			height := newHeights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			right := i
			res = max(res, (right-left-1)*height)
		}
		stack = append(stack, i)
	}

	return res
}
