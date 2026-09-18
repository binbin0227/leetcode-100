package main

import "math"

func myAtoi(s string) int {
	n := len(s)
	i := 0
	isPositive := 1
	var res int64

	// 1. 跳过空格
	for i < n && s[i] == ' ' {
		i++
	}
	if i >= n {
		return 0
	}

	// 2. 处理正负号
	if s[i] == '-' {
		isPositive = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	for i < n && s[i] >= '0' && s[i] <= '9' {
		curr := int(s[i] - '0')
		res = res*10 + int64(curr)

		// 实时检测有无溢出
		if isPositive == 1 && res > math.MaxInt32 {
			res = math.MaxInt32
		} else if isPositive == -1 && -res < math.MinInt32 {
			res = -math.MinInt32
		}

		i++
	}

	return isPositive * int(res)
}