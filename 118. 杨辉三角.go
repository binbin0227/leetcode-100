package main

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	oldNumRow := make([]int, 0)
	for i := range numRows {
		numRow := make([]int, 0, i+1)
		for j := range i + 1 {
			if j == 0 || j == i {
				numRow = append(numRow, 1)
				continue
			}
			numRow = append(numRow, oldNumRow[j-1]+oldNumRow[j])
		}
		res = append(res, numRow)
		oldNumRow = numRow[:]
	}
	return res
}
