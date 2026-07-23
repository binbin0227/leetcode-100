package main

func trap(height []int) int {
	// 先分别计算从最左手到（包括）当前位置和最右手（包括）到当前位置的最高水位
	// 然后遍历累加每格会积累的水量

	res := 0
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	for i := range n {
		if i == 0 {
			leftMax[0] = height[0]
			continue // 记得 continue
		}
		if leftMax[i-1] > height[i] {
			leftMax[i] = leftMax[i-1]
		} else {
			leftMax[i] = height[i]
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			rightMax[n-1] = height[n-1]
			continue
		}
		if rightMax[i+1] > height[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = height[i]
		}
	}

	for i := range n {
		res += max(0, min(leftMax[i], rightMax[i])-height[i])
	}
	return res
}
