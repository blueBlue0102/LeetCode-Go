package leetcode

func SpiralMatrixII(n int) [][]int {
	type Direction string
	const (
		Up    Direction = "上"
		Down  Direction = "下"
		Left  Direction = "左"
		Right Direction = "右"
	)
	var direction Direction = Right
	var x int = -1
	var y int = 0
	var result [][]int = make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for num := 1; num <= n*n; num++ {
		switch direction {
		case Right:
			x += 1
			if x+1 == n || result[y][x+1] != 0 {
				direction = Down
			}
		case Down:
			y += 1
			if y+1 == n || result[y+1][x] != 0 {
				direction = Left
			}
		case Left:
			x -= 1
			if x-1 == -1 || result[y][x-1] != 0 {
				direction = Up
			}
		case Up:
			y -= 1
			if y-1 == -1 || result[y-1][x] != 0 {
				direction = Right
			}
		}
		result[y][x] = num
	}
	return result
}
