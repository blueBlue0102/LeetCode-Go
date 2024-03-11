package leetcode

func SquaresofaSortedArray(nums []int) []int {
	var pointer = len(nums) - 1
	var leftIndex = 0
	var rightIndex = len(nums) - 1
	var result = make([]int, len(nums))
	for rightIndex >= leftIndex {
		if nums[leftIndex]*nums[leftIndex] >= nums[rightIndex]*nums[rightIndex] {
			result[pointer] = nums[leftIndex] * nums[leftIndex]
			leftIndex++
			pointer--
		} else {
			result[pointer] = nums[rightIndex] * nums[rightIndex]
			rightIndex--
			pointer--
		}
	}
	return result
}
