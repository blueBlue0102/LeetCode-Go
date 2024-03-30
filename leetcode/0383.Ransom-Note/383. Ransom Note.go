package leetcode

func RansomNote(ransomNote string, magazine string) bool {
	magazineDict := make(map[rune]int)

	for _, runeValue := range magazine {
		magazineDict[runeValue]++
	}

	for _, runeValue := range ransomNote {
		magazineDict[runeValue]--
		if magazineDict[runeValue] < 0 {
			return false
		}
	}
	return true
}
