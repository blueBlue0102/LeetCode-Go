package leetcode

func ClimbingStairs(n int) int {
	// return recur(n)
	return iterate(n)
}

func recur(n int) int {
	var data = map[int]int{}
	var recurInterval func(n int) int
	recurInterval = func(n int) int {
		if n == 1 {
			return 1
		} else if n == 2 {
			return 2
		}

		var nDown1 int
		var nDown2 int

		if val, existed := data[n-1]; existed {
			nDown1 = val
		} else {
			data[n-1] = ClimbingStairs(n - 1)
			nDown1 = data[n-1]
		}
		if val, existed := data[n-2]; existed {
			nDown2 = val
		} else {
			data[n-2] = ClimbingStairs(n - 2)
			nDown2 = data[n-2]
		}

		return nDown1 + nDown2
	}
	return recurInterval(n)
}

func iterate(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	v1, v2, v3 := 1, 2, 0

	for i := 2; i < n; i++ {
		v3 = v1 + v2
		v1, v2 = v2, v3
	}

	return v3
}
