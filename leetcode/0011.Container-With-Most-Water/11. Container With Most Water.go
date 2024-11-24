package leetcode

func ContainerWithMostWater(height []int) int {
	left, right := 0, len(height)-1
	maxCapacity := 0
	for right > left {
		capacity := min(height[left], height[right]) * (right - left)
		if capacity > maxCapacity {
			maxCapacity = capacity
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxCapacity
}
