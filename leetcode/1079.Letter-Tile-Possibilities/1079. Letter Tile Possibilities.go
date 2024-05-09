package leetcode

func LetterTilePossibilities(tiles string) int {
	total := 0
	// 紀錄每種字元的數量
	dict := map[rune]int{}
	for _, runeVal := range tiles {
		dict[runeVal]++
	}
	backtracking(&total, dict)
	return total
}

func backtracking(count *int, dict map[rune]int) {
	for runeVal, num := range dict {
		if num > 0 {
			dict[runeVal]--
			*count++
			backtracking(count, dict)
			dict[runeVal]++
		}
	}
}
