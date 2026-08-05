package main

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)

	isHuiwen := func(s string) bool {
		l, r := 0, len(s)-1
		for l <= r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if startIdx == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		for i := startIdx; i < len(s); i++ {
			if isHuiwen(s[startIdx : i+1]) {
				path = append(path, s[startIdx:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}
