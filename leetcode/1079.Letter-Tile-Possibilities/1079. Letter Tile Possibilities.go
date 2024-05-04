package leetcode

func LetterTilePossibilities(tiles string) int {
	ansMap := map[string]bool{}
	for length := 1; length <= len([]rune(tiles)); length++ {
		dfs(ansMap, []rune(tiles), []rune{}, length)
	}
	return len(ansMap)
}

func dfs(ansMap map[string]bool, tiles []rune, data []rune, length int) {
	if len(data) == length {
		if !ansMap[string(data)] {
			ansMap[string(data)] = true
		}
		return
	}
	for i, runeValue := range tiles {
		data = append(data, runeValue)
		dfs(ansMap, append(tiles[:i:i], tiles[i+1:]...), data, length)
		data = data[:len(data)-1]
	}
}
