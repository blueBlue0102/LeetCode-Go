package leetcode

func CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtracking(candidates, target, []int{}, 0, &result)
	return result
}

func backtracking(candidates []int, target int, selected []int, sum int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, selected...))
		return
	}
	for i := range candidates {
		if candidates[i]+sum <= target {
			selected = append(selected, candidates[i])
			backtracking(candidates[i:], target, selected, candidates[i]+sum, result)
			selected = selected[:len(selected)-1]
		}
	}
}
