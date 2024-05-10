package leetcode

import "strings"

func NQueens(n int) [][]string {
	result := [][]string{}
	backtracking(&result, n, 0, [][]int{})
	return result
}

// [[0,1], [1,0]] => [[".Q", "Q."]]
func intResult2StringResult(points [][]int) []string {
	n := len(points)
	ans := []string{}
	for _, point := range points {
		col := point[1]
		ans = append(ans, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
	}
	return ans
}

// 驗證該位置能否放置 queen
func verify(n int, points [][]int, x int, y int) bool {
	for _, point := range points {
		// same row?
		if point[0] == x {
			return false
		}
		// same col?
		if point[1] == y {
			return false
		}
		// same diagonal?
		// 右上右下都不用檢查，因為是從左邊開始放起
		// Upper Left
		for px, py := x, y; px >= 0 && py >= 0; px, py = px-1, py-1 {
			if px == point[0] && py == point[1] {
				return false
			}
		}
		// Lower Left
		for px, py := x, y; px >= 0 && py < n; px, py = px-1, py+1 {
			if px == point[0] && py == point[1] {
				return false
			}
		}
	}

	return true
}

func backtracking(result *[][]string, n int, row int, chessBoard [][]int) {
	if row == n {
		*result = append(*result, intResult2StringResult(chessBoard))
		return
	}
	for col := 0; col < n; col++ {
		if verify(n, chessBoard, row, col) {
			chessBoard = append(chessBoard, []int{row, col})
			backtracking(result, n, row+1, chessBoard)
			chessBoard = chessBoard[:len(chessBoard)-1]
		}
	}
}
