package leetcode

func SqrtX(x int) int {
	left := 0
	right := x
	for {
		k := (right-left)/2 + left
		if k*k <= x && (k+1)*(k+1) > x {
			return k
		} else if k*k > x {
			right = k - 1
		} else if k*k < x {
			left = k + 1
		}
	}
}
