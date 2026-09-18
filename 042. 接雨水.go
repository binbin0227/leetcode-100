package main

func trap(height []int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(height)
	tallestLeft, tallestRight := make([]int, n), make([]int, n)
	res := 0

	for i := 1; i < n; i++ {
		tallestLeft[i] = max(tallestLeft[i-1], height[i-1]) //不是 height[i]
	}
	for i := n - 2; i >= 0; i-- {
		tallestRight[i] = max(tallestRight[i+1], height[i+1])
	}

	for i := 0; i < n; i++ {
		res += max(0, min(tallestLeft[i], tallestRight[i])-height[i])
	}

	return res
}
