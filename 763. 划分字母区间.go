package main

func partitionLabels(s string) []int {
	// 先记录每个字符最后出现的位置
	res := make([]int, 0)
	lastPos := make(map[byte]int, 0)
	for i := range len(s) {
		lastPos[s[i]] = i
	}

	// 然后遍历 s ,如果这个字符出最后出现的位置在右指针右边，那就更新右指针
	// 等于右指针就切一刀
	l, r := 0, 0
	for i := range len(s) {
		if lastPos[s[i]] > r {
			r = lastPos[s[i]]
		}
		if i == r {
			res = append(res, r-l+1)
			l, r = i+1, i+1
		}
	}
	return res
}
