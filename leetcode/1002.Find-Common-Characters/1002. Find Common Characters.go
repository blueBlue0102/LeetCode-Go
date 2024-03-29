package leetcode

func FindCommonCharacters(words []string) []string {
	// 建立 word[0] 的 dict
	word0Dict := make(map[rune]int)
	for _, runeValue := range words[0] {
		word0Dict[runeValue]++
	}

	// 開始輪巡 word[1] 之後
	for _, word := range words[1:] {
		// 建立 word dict
		wordDict := make(map[rune]int)
		for _, runeValue := range word {
			wordDict[runeValue]++
		}

		// 與 word0 進行比較
		for runeValue, count := range word0Dict {
			word0Dict[runeValue] = min(count, wordDict[runeValue])
		}
	}

	// 放置結果至 result 中
	result := make([]string, 0)
	for runeValue, count := range word0Dict {
		for i := 0; i < count; i++ {
			result = append(result, string(runeValue))
		}
	}
	return result
}
