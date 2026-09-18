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



func longestPalindrome2(s string) string {
	check := func(leftIdx, rightIdx int) (int, int) {
		for leftIdx >= 0 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
			leftIdx--
			rightIdx++
		}

		return leftIdx + 1, rightIdx - 1
	}

	resL, resR := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := check(i, i)
		l2, r2 := check(i, i+1)

		l, r := 0, 0
		if r1-l1 > r2-l2 {
			l, r = l1, r1
		} else {
			l, r = l2, r2
		}

		if r-l > resR-resL {
			resL, resR = l, r
		}
	}

	return s[resL : resR+1]
}