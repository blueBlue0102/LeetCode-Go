package leetcode

func PascalTriangleII(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevLevelRow := PascalTriangleII(rowIndex - 1)
	result := make([]int, rowIndex+1)
	for index := range result {
		if index == 0 || index == rowIndex {
			result[index] = 1
		} else {
			result[index] = prevLevelRow[index-1] + prevLevelRow[index]
		}
	}
	return result
}
