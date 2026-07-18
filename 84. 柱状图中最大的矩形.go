package main

func largestRectangleArea(heights []int) int {
	newHeights := make([]int, len(heights)+2)
	copy(newHeights[1:], heights)

	stack := []int{}
	maxArea := 0

	for i := range len(newHeights) {
		for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
			curIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curHeight := newHeights[curIndex]

			leftIndex := stack[len(stack)-1]
			curWidth := i - leftIndex - 1

			area := curHeight * curWidth
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
