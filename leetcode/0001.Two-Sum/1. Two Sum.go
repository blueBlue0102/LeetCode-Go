package leetcode

func TwoSum(nums []int, target int) []int {
	// 建立字典 (target - 現在值): 值的 index
	var twoSumMap = make(map[int]int)
	// 開始掃 array
	for index := 0; index < len(nums); index++ {
		var neededValue = target - nums[index]
		if indexOfMapValue, ok := twoSumMap[nums[index]]; ok {
			return []int{indexOfMapValue, index}
		} else {
			twoSumMap[neededValue] = index
		}
	}
	return []int{0, 0}
}
