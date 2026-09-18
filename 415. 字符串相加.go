package main

func addStrings(num1 string, num2 string) string {
	res := make([]byte, 0)
	p1, p2 := len(num1)-1, len(num2)-1
	flag := 0

	for p1 >= 0 || p2 >= 0 || flag != 0 {
		var n1, n2 int
		if p1 >= 0 {
			n1 = int(num1[p1] - '0') // 转类型
			p1--
		}
		if p2 >= 0 {
			n2 = int(num2[p2] - '0') // 转类型
			p2--
		}

		sum := n1 + n2 + flag
		if sum >= 10 {
			flag = 1
			sum -= 10
		} else {
			flag = 0
		}

		res = append(res, byte(sum+'0')) // 转类型
	}

	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return string(res)
}