package utils

import (
	"cmp"
	"sort"
)

// 以 in-place 的方式排序二維陣列
func Sort2DArray[T cmp.Ordered](matrix [][]T) [][]T {
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
