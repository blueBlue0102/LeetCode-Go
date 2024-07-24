package leetcode

import "sort"

func SorttheJumbledNumbers(mapping []int, nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		n1, n2 := getMappedValue(mapping, nums[i]), getMappedValue(mapping, nums[j])
		if n1 == n2 {
			return i < j
		}
		return n1 < n2
	})
	return nums
}

func getMappedValue(mapping []int, num int) int {
	if num == 0 {
		return mapping[0]
	}

	result := 0
	base := 1
	for num > 0 {
		digit := mapping[num%10]
		result += digit * base
		base *= 10
		num /= 10
	}
	return result
}
