package main

func rand10() int {
	var a, b int
	for {
		a = rand7()
		b = rand7()
		if (a-1)*7+b <= 40 {
			break
		}
	}

	return ((a-1)*7+b)%10 + 1
}