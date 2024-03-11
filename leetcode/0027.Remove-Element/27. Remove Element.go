package leetcode

func RemoveElement(nums []int, val int) int {
	var k int
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}
