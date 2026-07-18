package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)
	for i := range n {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return res
}
