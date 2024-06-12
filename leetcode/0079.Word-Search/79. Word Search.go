package leetcode

type Point struct {
	row  int
	col  int
	used bool
}

func WordSearch(board [][]byte, word string) bool {
	// 建立 board map，以方便快速查詢特定字元的位置
	boardMap := map[byte][]Point{}
	for row := range board {
		for col := range board[row] {
			boardMap[board[row][col]] = append(boardMap[board[row][col]], Point{row, col, false})
		}
	}

	var backtracking func(wordBytes []byte, previousPoint Point, isFirstPoint bool) bool
	backtracking = func(wordBytes []byte, previousPoint Point, isFirstPoint bool) bool {
		if len(wordBytes) == 0 {
			return true
		}
		for pointIndex, point := range boardMap[wordBytes[0]] {
			if isFirstPoint || (!point.used && isAdjacentCell(point, previousPoint)) {
				if len(wordBytes) == 1 {
					return true
				}

				// backtracking
				boardMap[wordBytes[0]][pointIndex].used = true
				if backtracking(wordBytes[1:], point, false) {
					return true
				}
				boardMap[wordBytes[0]][pointIndex].used = false
			}
		}
		return false
	}
	return backtracking([]byte(word), Point{0, 0, false}, true)
}

func isAdjacentCell(p1 Point, p2 Point) bool {
	return ((p1.row == p2.row) && (p1.col == p2.col+1 || p1.col == p2.col-1)) || ((p1.col == p2.col) && (p1.row == p2.row+1 || p1.row == p2.row-1))
}
