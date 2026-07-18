package main

import "strings"

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	res := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			num = 0
			res = ""
		} else if ch == ']' {
			lastNum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			res = lastStr + strings.Repeat(res, lastNum)
		} else {
			res += string(ch)
		}
	}
	return res
}
