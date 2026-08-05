package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(startIdx, currSum int)
	backtrack = func(startIdx, currSum int) {
		if currSum > target {
			return
		}
		if currSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i := startIdx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, currSum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
