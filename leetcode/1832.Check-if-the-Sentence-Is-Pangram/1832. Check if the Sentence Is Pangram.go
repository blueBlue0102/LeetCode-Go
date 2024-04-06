package leetcode

func CheckiftheSentenceIsPangram(sentence string) bool {
	dict := make(map[rune]bool)
	for _, runeValue := range sentence {
		dict[runeValue] = true
		if len(dict) == 26 {
			return true
		}
	}
	return false
}
