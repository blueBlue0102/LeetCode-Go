package leetcode

import "math"

func KClosestPointstoOrigin(points [][]int, k int) [][]int {
	// 定義 pivot 以左都小於 points[pivot]；以右都大於等於 points[pivot]
	// points => { <= points[pivot] } + { > points[pivot] }
	pivot := len(points) - 1
	left := -1
	for left < pivot-1 {
		if CalDistance(points[pivot]) >= CalDistance(points[pivot-1]) {
			left++
			points[left], points[pivot-1] = points[pivot-1], points[left]
		} else {
			points[pivot], points[pivot-1] = points[pivot-1], points[pivot]
			pivot--
		}
	}

	if pivot > k-1 {
		return KClosestPointstoOrigin(points[:pivot], k)
	} else if pivot < k-1 {
		return append(points[:pivot], KClosestPointstoOrigin(points[pivot:], k-left-1)...)
	}
	return points[:k]
}

func CalDistance(point []int) float64 {
	return math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
}
