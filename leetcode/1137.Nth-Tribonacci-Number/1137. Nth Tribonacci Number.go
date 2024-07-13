package leetcode

func NthTribonacciNumber(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}

	var v3 int
	v0, v1, v2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		v3 = v0 + v1 + v2
		v0, v1, v2 = v1, v2, v3
	}
	return v3
}
