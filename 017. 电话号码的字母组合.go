package main

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	path := make([]byte, 0)

	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if len(path) == len(digits) {
			temp := make([]byte, len(path))
			copy(temp, path)
			res = append(res, string(temp))
		}

		if startIdx >= len(digits) {
			return
		}

		for _, word := range m[digits[startIdx]] {
			path = append(path, word)
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
