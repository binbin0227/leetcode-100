package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		if target > matrix[mid/n][mid%n] {
			left = mid + 1
		} else if target < matrix[mid/n][mid%n] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}
