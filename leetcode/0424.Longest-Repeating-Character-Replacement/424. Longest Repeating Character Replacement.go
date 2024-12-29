package leetcode

func LongestRepeatingCharacterReplacement(s string, k int) int {
	result := 0
	left := 0
	// 紀錄 sliding window 區間內字母的頻率
	letterFreq := map[byte]int{}
	mostLetterFreq := 0

	for right := 0; right < len(s); right++ {
		letterFreq[s[right]]++
		// 找出在這個區間下最多次數的字元
		mostLetterFreq = max(mostLetterFreq, letterFreq[s[right]])
		// 計算目前區間內使用了多少次 replace
		replaceCount := (right - left + 1) - mostLetterFreq
		// 如果超出 replace 的 quota，表示 left 需要往前
		// 由於每次 right 只會前進一步，所以若發生不夠，replace 只會缺 1，left 只需要前進一步
		if replaceCount == k+1 {
			letterFreq[s[left]]--
			left++
		}

		// 將這個區間的長度和最大值比較
		result = max(result, right-left+1)
	}
	return result
}
