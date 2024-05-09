package leetcode

func GenerateParentheses(n int) []string {
	result := []string{}
	backtracking(&result, n, n, "")
	return result
}

func backtracking(result *[]string, left int, right int, data string) {
	if left == 0 && right == 0 {
		*result = append(*result, data)
		return
	}
	if left > 0 {
		backtracking(result, left-1, right, data+"(")
	}
	if right > 0 && right > left {
		backtracking(result, left, right-1, data+")")
	}
}
