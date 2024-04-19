package leetcode

import (
	"testing"
)

// 填入 function input type
type parameters0324 struct {
	nums []int
}

func TestWiggleSortII(t *testing.T) {
	tests := []struct {
		parameters0324
	}{
		// 填入 test case
		{
			parameters0324{[]int{1, 5, 1, 1, 6, 4}},
		},
		{
			parameters0324{[]int{1, 3, 2, 2, 3, 1}},
		},
		{
			parameters0324{[]int{4, 5, 5, 6}},
		},
	}

	for _, test := range tests {
		t.Run("Test WiggleSortII", func(t *testing.T) {
			// 完整輸入參數
			WiggleSortII(test.parameters0324.nums)
			// compare 的方式需視情況調整
			if !verifyWiggleSort(test.parameters0324.nums) {
				t.Errorf("WiggleSortII(%+v) got %+v", test.parameters0324, test.parameters0324.nums)
			}
		})
	}
}

func verifyWiggleSort(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	prevDiff := nums[1] - nums[0]
	if prevDiff == 0 {
		return false
	}

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		// 判斷相鄰元素的差異是否不符合小於大於小於...的規律
		if prevDiff*diff >= 0 {
			return false
		}
		prevDiff = diff
	}

	return true
}
