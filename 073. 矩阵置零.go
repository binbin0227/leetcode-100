package main

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	rows := make([]int, 0)
	cols := make([]int, 0)

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	for _, i := range rows {
		for j := range n {
			matrix[i][j] = 0
		}
	}
	for _, j := range cols {
		for i := range m {
			matrix[i][j] = 0
		}
	}
}
