package leetcode

func IntersectionofTwoArrays(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	dict1 := make(map[int]bool)

	for _, num := range nums1 {
		dict1[num] = true
	}
	for _, num := range nums2 {
		if dict1[num] {
			result = append(result, num)
			delete(dict1, num)
		}
	}

	return result
}
