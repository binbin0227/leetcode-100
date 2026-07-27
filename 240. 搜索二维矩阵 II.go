package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1

	i := 0
	j := n

	for {
		if j < 0 || i > m {
			return false
		}

		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}
