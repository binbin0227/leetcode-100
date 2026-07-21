package main

import "sort"

func groupAnagrams(strs []string) [][]string {

	// 键为字母组成，值为符合要求的单词组成的切片
	m := make(map[string][]string)

	for _, str := range strs {

		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		key := string(chars)

		value, ok := m[key]
		if ok {
			// if 表中有同种组成的字母
			value = append(value, str)
			m[key] = value // 别忘记把 value 放回去
		} else {
			// if 表中没有同种组成的字母
			m[key] = []string{str}
		}
	}

	// 把所有值拎出来作为结果
	res := make([][]string, 0) // 一会要 append ，所以初始化长度为0
	for _, i := range m {
		res = append(res, i)
	}
	return res
}
