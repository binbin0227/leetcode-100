package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if len(path) == 2*n {
			temp := make([]byte, len(path))
			copy(temp, path)
			res = append(res, string(temp))
		}
		if len(path) >= 2*n {
			return
		}

		if left < n { // 左括号数量小于 n 就能放左括号
			path = append(path, '(')
			backtrack(left+1, right)
			path = path[:len(path)-1]
		}
		if left > right { // 左括号数量大于右括号数量，就能放右括号
			path = append(path, ')')
			backtrack(left, right+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
