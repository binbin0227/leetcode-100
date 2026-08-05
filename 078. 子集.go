package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		// 不用写 if startIdx >= len(nums) { return }
		// 如果越界了，for 循环不会执行

		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1) // 不是 startIdx + 1
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
