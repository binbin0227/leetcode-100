package main

func minWindow(s string, t string) string {
	res := ""
	tCount := make(map[byte]int, 0)
	for i := range len(t) {
		tCount[t[i]]++
	}

	wCount := make(map[byte]int, 0)
	l, r := 0, 0
	for r < len(s) {
		wCount[s[r]]++
		for check(wCount, tCount) {
			// If条件要写两个，否则永远返回空值
			if res == "" || len(s[l:r+1]) < len(res) {
				res = s[l : r+1]
			}
			wCount[s[l]]--
			l++
		}
		r++
	}
	return res
}

func check(wCount, tCount map[byte]int) bool {
	for tKey, tValue := range tCount {
		if wCount[tKey] < tValue {
			return false
		}
	}
	return true
}
