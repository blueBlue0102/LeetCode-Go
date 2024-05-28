package leetcode

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	seenChar := map[rune]int{}
	runeStr := []rune(s)
	longestLength := 0
	leftIndex := 0
	for rightIndex, char := range runeStr {
		repeatIndex, existed := seenChar[char]
		if existed {
			// move leftIndex to repeatIndex + 1
			// and remove dict member that between leftIndex and repeatIndex
			for leftIndex < repeatIndex+1 {
				delete(seenChar, runeStr[leftIndex])
				leftIndex++
			}
		} else {
			// renew longestLength
			longestLength = max(longestLength, rightIndex-leftIndex+1)
		}
		seenChar[char] = rightIndex
	}

	return longestLength
}
