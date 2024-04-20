package leetcode

func KthLargestElementinanArray(nums []int, k int) int {
	pivotVal := nums[len(nums)-1]
	left, curr, right := -1, 0, len(nums)
	// 將 nums 根據 pivotVal 由大排至小
	for curr < right {
		if nums[curr] > pivotVal {
			left++
			nums[curr], nums[left] = nums[left], nums[curr]
			curr++
		} else if nums[curr] < pivotVal {
			right--
			nums[curr], nums[right] = nums[right], nums[curr]
		} else if nums[curr] == pivotVal {
			curr++
		}
	}

	if k-1 <= left {
		// 找到的 pivotVal 太小了，所以要往 left 區間去找
		return KthLargestElementinanArray(nums[:left+1], k)
	} else if k-1 >= right {
		// 找到的 pivotVal 太大了，所以要往 right 區間去找
		return KthLargestElementinanArray(nums[right:], k-right)
	}
	// pivotVal 落在 left 和 right 之間
	// 由於 left 和 right 之間的值都必定等於 pivotVal，所以回傳 pivotVal
	return pivotVal
}
