package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	u, d, l, r := 0, m-1, 0, n-1
	res := make([]int, 0, m*n)

	for {
		if l > r {
			break
		}
		for j := l; j <= r; j++ {
			res = append(res, matrix[u][j])
		}
		u++

		if u > d {
			break
		}
		for i := u; i <= d; i++ {
			res = append(res, matrix[i][r])
		}
		r--

		if l > r {
			break
		}
		for j := r; j >= l; j-- {
			res = append(res, matrix[d][j])
		}
		d--

		if u > d {
			break
		}
		if l <= r {
			for i := d; i >= u; i-- {
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}
