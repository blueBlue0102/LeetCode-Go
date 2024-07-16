package leetcode

func HouseRobberII(nums []int) int {
	return max(nums[0]+dp(nums, 2, len(nums)-2), dp(nums, 1, len(nums)-1))
}

func dp(nums []int, startIndex int, endIndex int) int {
	length := endIndex - startIndex + 1
	if length <= 0 {
		return 0
	} else if length == 1 {
		return nums[startIndex]
	}

	nums = nums[startIndex : endIndex+1]
	v0, v1 := nums[0], max(nums[0], nums[1])
	v2 := v1
	for i := 2; i < len(nums); i++ {
		v2 = max(nums[i]+v0, v1)
		v0, v1 = v1, v2
	}
	return v2
}
