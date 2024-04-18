package leetcode

func KthLargestElementinanArray(nums []int, k int) int {
	left, pivot := -1, len(nums)-1
	// 最終 pivot 的右邊都是 > nums[pivot]；左邊都是 <= nums[pivot]
	for left < pivot-1 {
		if nums[pivot] >= nums[pivot-1] {
			left++
			nums[left], nums[pivot-1] = nums[pivot-1], nums[left]
		} else {
			nums[pivot], nums[pivot-1] = nums[pivot-1], nums[pivot]
			pivot--
		}
	}

	if pivot > len(nums)-k {
		// 所找到的 pivot 太大，下一次從 pivot 的左邊開始找起
		return KthLargestElementinanArray(nums[:pivot], k-(len(nums)-pivot))
	} else if pivot < len(nums)-k {
		// 所找到的 pivot 太小，下一次從 pivot 的右邊開始找起
		return KthLargestElementinanArray(nums[pivot+1:], k)
	}
	return nums[pivot]
}
