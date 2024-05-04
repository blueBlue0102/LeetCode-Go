package leetcode

import (
	"unicode"
)

func ValidPalindrome(s string) bool {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left <= right {
		leftRune := unicode.ToLower(runes[left])
		if !isAlphanumeric(leftRune) {
			left++
			continue
		}
		rightRune := unicode.ToLower(runes[right])
		if !isAlphanumeric(rightRune) {
			right--
			continue
		}

		if leftRune != rightRune {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z')
}
