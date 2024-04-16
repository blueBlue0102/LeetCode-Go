package utils

import "sort"

// 以 in-place 的方式排序二維 int 陣列
func Sort2DIntArray(matrix [][]int) [][]int {
	sort.SliceStable(matrix, func(i, j int) bool {
		minLen := min(len(matrix[i]), len(matrix[j]))
		for k := 0; k < minLen; k++ {
			if matrix[i][k] != matrix[j][k] {
				return matrix[i][k] < matrix[j][k]
			}
		}
		return len(matrix[i]) < len(matrix[j])
	})
	return matrix
}
