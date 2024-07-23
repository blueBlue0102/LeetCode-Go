package leetcode

import "sort"

func SortArraybyIncreasingFrequency(nums []int) []int {
	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num]++
	}
	sort.Slice(nums, func(a, b int) bool {
		if freqMap[nums[a]] == freqMap[nums[b]] {
			return nums[a] > nums[b]
		}
		return freqMap[nums[a]] < freqMap[nums[b]]
	})
	return nums
}
