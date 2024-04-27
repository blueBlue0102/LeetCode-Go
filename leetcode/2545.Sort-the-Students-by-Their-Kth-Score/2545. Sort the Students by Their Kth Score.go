package leetcode

import (
	"math/rand/v2"
)

func SorttheStudentsbyTheirKthScore(score [][]int, k int) [][]int {
	return QuickSort(score, k)
}

func QuickSort(score [][]int, k int) [][]int {
	if len(score) <= 1 {
		return score
	}

	left, right := 0, len(score)-1

	// Pick score pivot
	// Move the pivot to the right
	pivotIndex := rand.Int() % len(score)
	score[pivotIndex], score[right] = score[right], score[pivotIndex]

	// Pile elements bigger than the pivot on the left
	for curr := range score {
		if score[curr][k] > score[right][k] {
			score[curr], score[left] = score[left], score[curr]
			left++
		}
	}
	// Place the pivot after the last bigger element
	score[left], score[right] = score[right], score[left]

	QuickSort(score[:left], k)
	QuickSort(score[left+1:], k)

	return score
}
