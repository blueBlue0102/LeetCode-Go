package leetcode

func DeleteandEarn(nums []int) int {
	points := map[int]int{}
	maxNumInNums := 0
	for _, num := range nums {
		points[num] += num
		maxNumInNums = max(maxNumInNums, num)
	}

	// 從 0 逐漸遞增到 maxNumInNums 為止
	v0, v1 := 0, points[1]
	v2 := v1
	for i := 2; i <= maxNumInNums; i++ {
		v2 = max(points[i]+v0, v1)
		v0, v1 = v1, v2
	}
	return v2
}
