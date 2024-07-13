package leetcode

func MinCostClimbingStairs(cost []int) int {
	// 已知 len(cost) >= 2
	minCosts := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		minCosts[i] = min(minCosts[i-1]+cost[i-1], minCosts[i-2]+cost[i-2])
	}
	return minCosts[len(cost)]
}
