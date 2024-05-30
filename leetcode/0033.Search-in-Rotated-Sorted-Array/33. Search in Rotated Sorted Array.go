package leetcode

func SearchinRotatedSortedArray(nums []int, target int) int {
	// `[1 2 3 4 5 6 7]`
	// `[7 1 2 3 4 5 6]`
	// `[6 7 1 2 3 4 5]`
	// `[5 6 7 1 2 3 4]`
	// `[4 5 6 7 1 2 3]`
	// `[3 4 5 6 7 1 2]`
	// `[2 3 4 5 6 7 1]`
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] {
			// `[1 2 3 4 5 6 7]`
			// `[7 1 2 3 4 5 6]`
			// `[6 7 1 2 3 4 5]`
			// `[5 6 7 1 2 3 4]`
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				if target > nums[right] {
					right = mid - 1
				} else if target < nums[right] {
					left = mid + 1
				} else {
					return right
				}
			}
		} else if nums[mid] > nums[right] {
			// `[4 5 6 7 1 2 3]`
			// `[3 4 5 6 7 1 2]`
			// `[2 3 4 5 6 7 1]`
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				if target < nums[left] {
					left = mid + 1
				} else if target > nums[left] {
					right = mid - 1
				} else {
					return left
				}
			}
		}
	}

	// 此時必定 left == right
	if nums[left] == target {
		return left
	} else {
		return -1
	}
}
