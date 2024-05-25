package leetcode

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	recur(image, sr, sc, image[sr][sc], color)
	return image
}

func recur(image [][]int, r int, c int, startColor int, endColor int) {
	image[r][c] = endColor
	if r > 0 && image[r-1][c] == startColor {
		recur(image, r-1, c, startColor, endColor)
	}
	if c > 0 && image[r][c-1] == startColor {
		recur(image, r, c-1, startColor, endColor)
	}
	if r < len(image)-1 && image[r+1][c] == startColor {
		recur(image, r+1, c, startColor, endColor)
	}
	if c < len(image[0])-1 && image[r][c+1] == startColor {
		recur(image, r, c+1, startColor, endColor)
	}
}
