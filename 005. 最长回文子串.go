package main

func longestPalindrome(s string) string {
	res := ""
	for i := range len(s) {
		// 回文子串的长度可能是奇数，也可能是偶数
		length := max(check(s, i, i), check(s, i, i+1))
		if length > len(res) {
			res = s[i-(length-1)/2 : i+length/2+1]
		}
	}
	return res
}

func check(s string, l, r int) int {
	for 0 <= l && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return r - l - 1 // 回文子串长度
}
