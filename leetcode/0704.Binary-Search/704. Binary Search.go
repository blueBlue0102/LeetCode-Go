package leetcode

func BinarySearch(nums []int, target int) int {
	var leftIndex int = 0
	var rightIndex int = len(nums) - 1
	var targetIndex int
	for leftIndex <= rightIndex {
		targetIndex = (rightIndex-leftIndex)/2 + leftIndex
		if nums[targetIndex] == target {
			return targetIndex
		} else if nums[targetIndex] > target {
			rightIndex = targetIndex - 1
		} else if nums[targetIndex] < target {
			leftIndex = targetIndex + 1
		}
	}
	return -1
}
