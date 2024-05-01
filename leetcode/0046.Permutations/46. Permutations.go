package leetcode

func Permutations(nums []int) [][]int {
	ans := [][]int{}
	Enumerate(&ans, nums, []int{})
	return ans
}

func Enumerate(ans *[][]int, nums []int, data []int) {
	if len(nums) <= 1 {
		*ans = append(*ans, append(data, nums...))
		return
	}
	for i := range nums {
		// nums[:i:i] 是 Full slice expressions 的寫法
		// 為了使 `append` 必定 allocate new array，進而避免異動到 `nums`
		Enumerate(ans, append(nums[:i:i], nums[i+1:]...), append(data, nums[i]))
	}
}
