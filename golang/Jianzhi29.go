package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	top, bottom, left, right, ret := 0, len(matrix)-1, 0, len(matrix[0])-1, []int{}

	for true {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top += 1
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right -= 1
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			ret = append(ret, matrix[bottom][i])
		}
		bottom -= 1
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			ret = append(ret, matrix[i][left])
		}
		left += 1
		if left > right {
			break
		}
	}
	return ret
}
