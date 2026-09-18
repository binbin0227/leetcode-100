package main

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)

	for i := range len(s) {
		if j, ok := m[s[i]]; ok {
			// s[i] 是右括号
			if len(stack) == 0 || stack[len(stack)-1] != j { // 防越界
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// s[i] 是左括号
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
