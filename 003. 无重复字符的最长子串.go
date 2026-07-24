package main

func lengthOfLongestSubstring(s string) int {
	// 右指针从左往右扫,扫到的字符都放到滑动窗口里,
	// 如果扫到的字符已经在窗口里了，就删除滑动窗口里的这个字符并不断右移左指针直到无重复

	if len(s) == 0 {
		return 0
	}

	res := 0
	l := 0
	r := 0
	window := make(map[byte]struct{}, 0)

	for r < len(s) {
		if _, exist := window[s[r]]; exist {
			delete(window, s[l])
			l++
		} else {
			window[s[r]] = struct{}{}
			r++
			res = max(res, r-l)
		}
	}

	return res
}
