package main

func findAnagrams(s string, p string) []int {
	// 用两个长度为26的数组来记录 p 和滑动窗口的组成

	res := make([]int, 0)
	if len(p) > len(s) { // 特殊值处理
		return res
	}

	pCount := [26]int{}
	for _, char := range []byte(p) {
		pCount[char-'a']++
	}

	wCount := [26]int{}
	l := 0
	r := len(p) - 1
	for i := range r + 1 {
		wCount[s[i]-'a']++
	}

	for r < len(s) {
		if wCount == pCount {
			res = append(res, l)
		}

		if r == len(s)-1 {
			break
		}

		r++
		wCount[s[r]-'a']++
		wCount[s[l]-'a']--
		l++
	}

	return res
}
