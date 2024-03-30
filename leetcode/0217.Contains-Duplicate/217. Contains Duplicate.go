package leetcode

func ContainsDuplicate(nums []int) bool {
	dict := make(map[int]bool)
	for _, n := range nums {
		if dict[n] {
			return true
		} else {
			dict[n] = true
		}
	}
	return false
}
