package main

import "strings"

func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0)
	num := 0
	str := ""

	for _, word := range s {
		if word >= '0' && word <= '9' {
			num = num*10 + int(word-'0') // 这里的word是 ASCII 码
		} else if word >= 'a' && word <= 'z' {
			str = str + string(word)
		} else if word == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num = 0
			str = ""
		} else {
			lastNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			str = lastStr + strings.Repeat(str, lastNum)
		}
	}

	return str
}
