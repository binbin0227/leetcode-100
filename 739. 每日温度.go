package main

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))

	for i := range len(temperatures) {
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			res[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
