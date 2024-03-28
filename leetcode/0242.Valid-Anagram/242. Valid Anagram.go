package leetcode

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dictMap := make(map[rune]int)
	for _, runeValue := range s {
		dictMap[runeValue]++
	}
	for _, runeValue := range t {
		dictMap[runeValue]--
	}

	for _, value := range dictMap {
		if value != 0 {
			return false
		}
	}
	return true
}
