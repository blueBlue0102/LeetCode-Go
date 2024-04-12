package leetcode

var data = make(map[int]int)

func ClimbingStairs(n int) int {
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
