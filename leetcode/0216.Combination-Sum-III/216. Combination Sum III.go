package leetcode

func CombinationSumIII(k int, n int) [][]int {
	result := [][]int{}
	backtracking(&result, k, n, 1, []int{})
	return result
}

func backtracking(result *[][]int, k int, n int, initVal int, data []int) {
	if k == 0 && n == 0 {
		// 必須建立新的 Slice，否則 data 會持續被異動
		newData := make([]int, len(data))
		copy(newData, data)
		*result = append(*result, newData)
		return
	}

	for val := initVal; val <= 9; val++ {
		if val > n {
			break
		}
		data = append(data, val)
		backtracking(result, k-1, n-val, val+1, data)
		data = data[:len(data)-1]
	}
}
