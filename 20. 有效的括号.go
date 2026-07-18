package main

func isValid(s string) bool {
	// 奇数直接返回
	if len(s)%2 != 0 {
		return false
	}

	stack := []byte{}
	mapping := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := range len(s) {
		char := s[i]
		if val, ok := mapping[char]; ok {
			// A. 左括号
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false // 栈空或括号不匹配
			}
			stack = stack[:len(stack)-1] // 弹出栈顶
		} else {
			// B. 右括号
			stack = append(stack, char)
		}
	}
	return len(stack) == 0 // 栈空才说明全匹配
}
