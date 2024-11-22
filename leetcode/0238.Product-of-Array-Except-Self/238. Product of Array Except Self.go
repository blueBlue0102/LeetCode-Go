package leetcode

func ProductofArrayExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}

	prefixProduct, suffixProduct := 1, 1
	for i := 0; i < len(nums); i++ {
		// prefix
		ans[i] *= prefixProduct
		prefixProduct *= nums[i]

		// suffix
		ans[len(nums)-i-1] *= suffixProduct
		suffixProduct *= nums[len(nums)-i-1]
	}
	return ans
}
