package leetcode

func WiggleSortII(nums []int) {
	// 取得 nums 中的中位數
	medianValue := quickSelect(nums, len(nums)/2)

	wiggleIndex := func(index int) int {
		mid := len(nums) / 2
		if index < mid {
			return 2*(index) + 1
		} else {
			return 2 * (index - mid)
		}
	}

	// 使用中位數對 nums 進行 DNF 排序，由大至小
	left, curr, right := -1, 0, len(nums)
	for curr < right {
		if nums[wiggleIndex(curr)] > medianValue {
			left++
			nums[wiggleIndex(curr)], nums[wiggleIndex(left)] = nums[wiggleIndex(left)], nums[wiggleIndex(curr)]
			curr++
		} else if nums[wiggleIndex(curr)] < medianValue {
			right--
			nums[wiggleIndex(curr)], nums[wiggleIndex(right)] = nums[wiggleIndex(right)], nums[wiggleIndex(curr)]
		} else {
			curr++
		}
	}
}

// 找出 k_th smallest 的 value，並且排序使 nums 的順序變成 => [<= nums[k]] + [> nums[k]]
// target: 0~len(nums)-1
func quickSelect(nums []int, target int) int {
	left, pivot := -1, len(nums)-1
	for left < pivot-1 {
		if nums[pivot] >= nums[pivot-1] {
			left++
			nums[left], nums[pivot-1] = nums[pivot-1], nums[left]
		} else {
			nums[pivot], nums[pivot-1] = nums[pivot-1], nums[pivot]
			pivot--
		}
	}

	if pivot < target {
		return quickSelect(nums[pivot+1:], target-pivot-1)
	} else if pivot > target {
		return quickSelect(nums[:pivot], target)
	}
	return nums[pivot]
}
