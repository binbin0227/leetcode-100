package main

func rotate(matrix [][]int) {

	//      0 1 2 3 4
	//    0 A A A B B
	//    1 A A A B B
	//    2 D D * B B
	//    3 D D C C C
	//    4 D D C C C

	N := len(matrix) // 阶数

	// 巧妙地统一了阶数为奇或者偶的情况
	for i := range N / 2 {
		for j := range (N + 1) / 2 {
			matrix[i][j],
				matrix[j][N-i-1],
				matrix[N-i-1][N-j-1],
				matrix[N-j-1][i] = matrix[N-j-1][i],
				matrix[i][j],
				matrix[j][N-i-1],
				matrix[N-i-1][N-j-1]
		}
	}
}
