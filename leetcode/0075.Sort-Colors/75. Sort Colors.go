package leetcode

func SortColors(nums []int) {
	const MIN, MAX = 0, 2
	left, mid, right := -1, 0, len(nums)
	for mid < right {
		if nums[mid] == MIN {
			left++
			nums[mid], nums[left] = nums[left], nums[mid]
			mid++
		} else if nums[mid] == MAX {
			right--
			nums[mid], nums[right] = nums[right], nums[mid]
		} else {
			mid++
		}
	}
}
