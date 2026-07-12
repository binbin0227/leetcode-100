package main

func partition(s string) [][]string {
	var res [][]string
	isPalindrome := func(left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	var backtrack func(path []string, start_idx int)
	backtrack = func(path []string, start_idx int) {
		if start_idx == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start_idx; i < len(s); i++ {
			if isPalindrome(start_idx, i) {
				curr := s[start_idx : i+1]
				path = append(path, curr)
				backtrack(path, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack([]string{}, 0)
	return res
}
