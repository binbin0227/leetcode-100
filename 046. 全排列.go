package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

	outer:
		for i := range len(nums) {
			for j := range len(path) {
				if nums[i] == path[j] {
					continue outer
				}
			}

			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}
