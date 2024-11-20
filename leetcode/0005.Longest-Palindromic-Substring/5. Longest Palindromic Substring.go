package leetcode

func LongestPalindromicSubstring(s string) string {
	longestLen := 0
	longestString := ""
	runeSlice := []rune(s)
	for index := range runeSlice {
		// odd mode
		length := 1
		leftIndex := index - 1
		rightIndex := index + 1
		for leftIndex >= 0 && rightIndex < len(s) {
			if runeSlice[leftIndex] == runeSlice[rightIndex] {
				length += 2
				leftIndex--
				rightIndex++
			} else {
				break
			}
		}
		if length > longestLen {
			longestLen = length
			longestString = string(runeSlice[leftIndex+1 : rightIndex])
		}
		// even mode
		length = 0
		leftIndex = index
		rightIndex = index + 1
		for leftIndex >= 0 && rightIndex < len(s) {
			if runeSlice[leftIndex] == runeSlice[rightIndex] {
				length += 2
				leftIndex--
				rightIndex++
			} else {
				break
			}
		}

		if length > longestLen {
			longestLen = length
			longestString = string(runeSlice[leftIndex+1 : rightIndex])
		}
	}

	return longestString
}
