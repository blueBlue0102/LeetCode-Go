package leetcode

func HouseRobber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	v0, v1 := nums[0], max(nums[0], nums[1])
	v2 := v1

	for i := 2; i < len(nums); i++ {
		v2 = max(v1, nums[i]+v0)
		v0, v1 = v1, v2
	}

	return v2

	// -------- array version
	// arr := make([]int, len(nums))
	// arr[0] = nums[0]
	// arr[1] = max(nums[0], nums[1])

	// for i := 2; i < len(nums); i++ {
	// 	arr[i] = max(arr[i-1], nums[i]+arr[i-2])
	// }

	// return arr[len(nums)-1]
}
