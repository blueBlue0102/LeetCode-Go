package leetcode

func SortColors(nums []int) {
	const MIN = 0
	const MAX = 2

	left := -1
	mid := 0
	right := len(nums)
	for mid < right {
		if nums[mid] == MIN {
			left++
			n := nums[mid]
			nums[mid] = nums[left]
			nums[left] = n

			mid++
		} else if nums[mid] == MAX {
			right--
			n := nums[mid]
			nums[mid] = nums[right]
			nums[right] = n
		} else {
			mid++
		}

	}
}
