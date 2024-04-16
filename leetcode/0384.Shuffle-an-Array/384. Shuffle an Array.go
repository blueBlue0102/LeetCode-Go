package leetcode

import (
	"math/rand/v2"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (solution *Solution) Reset() []int {
	return solution.nums
}

func (solution *Solution) Shuffle() []int {
	result := make([]int, len(solution.nums))
	copy(result, solution.nums)

	for index := len(result) - 1; index > 0; index-- {
		randNum := rand.IntN(index)
		result[randNum], result[index] = result[index], result[randNum]
	}
	return result
}
