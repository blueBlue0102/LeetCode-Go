package leetcode

import (
	"math/rand/v2"
)

func SortanArray(nums []int) []int {
	// return MergeSort_Recur(nums)
	return QuickSort_3Way(nums)
}

func MergeSort_Recur(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// Divide
	mid := len(nums) / 2
	leftSortedNums := MergeSort_Recur(nums[:mid])
	rightSortedNums := MergeSort_Recur(nums[mid:])

	// Merge
	return merge(leftSortedNums, rightSortedNums)
}

func merge(arr1 []int, arr2 []int) []int {
	p1, p2 := 0, 0
	result := []int{}

	// sort
	for p1 < len(arr1) && p2 < len(arr2) {
		if arr1[p1] <= arr2[p2] {
			result = append(result, arr1[p1])
			p1++
		} else {
			result = append(result, arr2[p2])
			p2++
		}
	}

	// 接上不需要排序的後段
	if p1 < len(arr1) {
		result = append(result, arr1[p1:]...)
	}
	if p2 < len(arr2) {
		result = append(result, arr2[p2:]...)
	}

	return result
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 最終 nums[< left] 的值都是小於 pivot 的值
	left, right := 0, len(nums)-1
	pivot := rand.Int() % len(nums)
	nums[pivot], nums[right] = nums[right], nums[pivot]

	for curr := range nums {
		if nums[curr] < nums[right] {
			nums[curr], nums[left] = nums[left], nums[curr]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]

	QuickSort(nums[:left])
	QuickSort(nums[left+1:])

	return nums
}

func QuickSort_3Way(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivotValue := nums[rand.Int()%len(nums)]
	left, curr, right := -1, 0, len(nums)

	for curr < right {
		if nums[curr] < pivotValue {
			left++
			nums[curr], nums[left] = nums[left], nums[curr]
			curr++
		} else if nums[curr] > pivotValue {
			right--
			nums[curr], nums[right] = nums[right], nums[curr]
		} else {
			curr++
		}
	}

	QuickSort_3Way(nums[:left+1])
	QuickSort_3Way(nums[right:])

	return nums
}
