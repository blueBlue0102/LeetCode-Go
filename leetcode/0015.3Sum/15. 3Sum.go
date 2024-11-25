package leetcode

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	// sort nums: O(nlogn)
	slices.Sort(nums)

	// find a, b, c that nums[a]+nums[b]+nums[c]=0: O(n^2)
	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b := a + 1
		c := len(nums) - 1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum > 0 {
				num := nums[c]
				for c > 0 && num == nums[c] {
					c--
				}
			} else if sum < 0 {
				num := nums[b]
				for b < len(nums) && num == nums[b] {
					b++
				}
			} else {
				result = append(result, []int{nums[a], nums[b], nums[c]})
				num := nums[b]
				for b < len(nums) && num == nums[b] {
					b++
				}
			}
		}
	}
	return result
}
