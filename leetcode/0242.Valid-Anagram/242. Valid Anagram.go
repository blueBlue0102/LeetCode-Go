package leetcode

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := map[rune]int{}

	for _, r := range s {
		dict[r]++
	}
	for _, r := range t {
		dict[r]--
		if dict[r] == 0 {
			delete(dict, r)
		}
	}

	return len(dict) == 0
}
