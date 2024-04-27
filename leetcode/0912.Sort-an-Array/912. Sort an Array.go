package leetcode

func SortanArray(nums []int) []int {
	return MergeSort_Recur(nums)
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
