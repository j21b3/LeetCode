package main

/*
*矩阵从左往右，从上到下各自递增，因此从(0,len(matrix[0])-1)开始可以做二分查找，类似二分树
 */

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) <= 0 { //empty matrix
		return false
	}

	i, j, l := 0, len(matrix[0])-1, len(matrix)

	for i < l && j >= 0 {
		if matrix[i][j] == target {
			return true //already found
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}
	}
	return false // not found
}
