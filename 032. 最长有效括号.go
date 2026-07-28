package main

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	// F(i) = 2 + F(i-2)
	// F(i) = 2 + F(i-1) + F(i-F(i-1)-2)
	dp := make([]int, len(s))
	dp[0] = 0

	// s[i]=='(' -> continue
	//       ')' -> s[i-1]=='(' -> dp[i]=dp[i-2]+2
	// 	                    ')' -> s[i-dp[i-1]-1]=='(' -> dp[i]=dp[i-1]+2+dp[i-dp[i-1]-2]

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		} else { // s[i]== ')'
			if s[i-1] == '(' {
				if i == 1 { // 防越界
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else { // s[i-1] == ')'
				if i-dp[i-1]-2 == -1 {
					if s[i-dp[i-1]-1] == '(' {
						dp[i] = dp[i-1] + 2
					}
				}
				if i-dp[i-1]-2 >= 0 {
					if s[i-dp[i-1]-1] == '(' {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					}
				}
			}
		}
	}

	res := 0
	for _, count := range dp {
		res = max(res, count)
	}
	return res
}
