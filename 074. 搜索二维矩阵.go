package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		if target < matrix[row][col] {
			right = mid - 1
		} else if target > matrix[row][col] {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}