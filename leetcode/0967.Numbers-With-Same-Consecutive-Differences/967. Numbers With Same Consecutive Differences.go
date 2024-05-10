package leetcode

func NumbersWithSameConsecutiveDifferences(n int, k int) []int {
	result := []int{}
	for firstDigit := 1; firstDigit <= 9; firstDigit++ {
		backtracking(&result, n-1, k, firstDigit, firstDigit)
	}
	return result
}

func backtracking(result *[]int, n int, k int, lastDigit int, total int) {
	if n == 0 {
		*result = append(*result, total)
		return
	}
	if lastDigit+k <= 9 {
		backtracking(result, n-1, k, lastDigit+k, total*10+(lastDigit+k))
	}
	if k != 0 && lastDigit-k >= 0 {
		backtracking(result, n-1, k, lastDigit-k, total*10+(lastDigit-k))
	}
}
