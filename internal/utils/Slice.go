package utils

import (
	"cmp"
	"sort"
)

// Sort2DArray sorts a 2D array by comparing elements in each subarray lexicographically.
// If the elements are equal up to the length of the shorter subarray,
// the shorter subarray is considered smaller.
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

// Sort2DArrayWithSortedSubArrays sorts a 2D array where each subarray is first sorted,
// then the 2D array is sorted based on the sorted subarrays
func Sort2DArrayWithSortedSubArrays[T cmp.Ordered](matrix [][]T) [][]T {
	// First sort each subarray
	for i := range matrix {
		sort.SliceStable(matrix[i], func(a, b int) bool {
			return matrix[i][a] < matrix[i][b]
		})
	}

	// Then sort the 2D array based on sorted subarrays
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
