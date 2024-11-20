package leetcode

func LongestConsecutiveSequence(nums []int) int {
	numsDict := map[int]bool{}
	for _, num := range nums {
		numsDict[num] = true
	}

	longestDist := 0
	for num, exist := range numsDict {
		if !exist {
			continue
		}
		dist := 1
		delete(numsDict, num)
		biggerCount := 0
		for numsDict[num+biggerCount+1] {
			delete(numsDict, num+biggerCount+1)
			biggerCount++
		}
		smallerCount := 0
		for numsDict[num-smallerCount-1] {
			delete(numsDict, num-smallerCount-1)
			smallerCount++
		}
		dist = dist + smallerCount + biggerCount

		if dist > longestDist {
			longestDist = dist
		}
	}

	return longestDist
}
