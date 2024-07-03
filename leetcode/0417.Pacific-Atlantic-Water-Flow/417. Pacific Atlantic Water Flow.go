package leetcode

type Coordinate struct {
	row int
	col int
}

func PacificAtlanticWaterFlow(heights [][]int) [][]int {
	result := [][]int{}
	if len(heights) == 0 {
		return result
	}
	MAP_ROW, MAP_COL := len(heights), len(heights[0])
	pacificMap, atlanticMap := map[Coordinate]bool{}, map[Coordinate]bool{}

	var explore func(target Coordinate, flowMap map[Coordinate]bool, prevHeight int)
	explore = func(target Coordinate, flowMap map[Coordinate]bool, prevHeight int) {
		// target 是否已巡過
		if flowMap[target] ||
			// target 是否超出 map
			(target.row < 0 || target.row == MAP_ROW || target.col < 0 || target.col == MAP_COL) ||
			// target 是否比前一個點更高或相等
			(heights[target.row][target.col] < prevHeight) {
			return
		}
		flowMap[target] = true
		explore(Coordinate{row: target.row - 1, col: target.col}, flowMap, heights[target.row][target.col])
		explore(Coordinate{row: target.row + 1, col: target.col}, flowMap, heights[target.row][target.col])
		explore(Coordinate{row: target.row, col: target.col - 1}, flowMap, heights[target.row][target.col])
		explore(Coordinate{row: target.row, col: target.col + 1}, flowMap, heights[target.row][target.col])

	}

	// pacific 上至下；atlantic 下至上
	for col := 0; col < MAP_COL; col++ {
		explore(Coordinate{row: 0, col: col}, pacificMap, heights[0][col])
		explore(Coordinate{row: MAP_ROW - 1, col: col}, atlanticMap, heights[MAP_ROW-1][col])
	}

	// pacific 右至左；atlantic 左至右
	for row := 0; row < MAP_ROW; row++ {
		explore(Coordinate{row: row, col: 0}, pacificMap, heights[row][0])
		explore(Coordinate{row: row, col: MAP_COL - 1}, atlanticMap, heights[row][MAP_COL-1])
	}

	for coordinate, visit := range pacificMap {
		if visit && atlanticMap[coordinate] {
			result = append(result, []int{coordinate.row, coordinate.col})
		}
	}
	return result
}
