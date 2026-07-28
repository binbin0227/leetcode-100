package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	p2, p1 := 1, 2
	var res int
	for i := 3; i <= n; i++ {
		res = p1 + p2
		p2, p1 = p1, res
	}

	return res
}
