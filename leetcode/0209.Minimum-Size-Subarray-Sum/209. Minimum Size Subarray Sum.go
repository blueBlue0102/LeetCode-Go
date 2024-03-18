package leetcode

func MinimumSizeSubarraySum(target int, nums []int) int {
	leftIndex := 0
	windowSum := 0
	const IMPOSSIBLE_VALUE = 999999
	minSize := IMPOSSIBLE_VALUE

	for rightIndex := 0; rightIndex < len(nums); rightIndex++ {
		windowSum += nums[rightIndex]
		for windowSum >= target {
			minSize = min(minSize, rightIndex-leftIndex+1)
			windowSum -= nums[leftIndex]
			leftIndex++
		}
	}

	if minSize == IMPOSSIBLE_VALUE {
		return 0
	} else {
		return minSize
	}
}
