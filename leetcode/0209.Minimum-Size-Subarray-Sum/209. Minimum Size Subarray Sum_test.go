package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0209 struct {
	target int
	nums   []int
}

func TestMinimumSizeSubarraySum(t *testing.T) {
	tests := []struct {
		parameters0209
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0209{target: 7, nums: []int{2, 3, 1, 2, 4, 3}},
			2,
		},
		{
			parameters0209{target: 4, nums: []int{1, 4, 4}},
			1,
		},
		{
			parameters0209{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			0,
		},
	}

	for _, test := range tests {
		t.Run("Test MinimumSizeSubarraySum", func(t *testing.T) {
			// 完整輸入參數
			result := MinimumSizeSubarraySum(test.parameters0209.target, test.parameters0209.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("MinimumSizeSubarraySum(%+v) got %+v, want %+v", test.parameters0209, result, test.ans)
			}
		})
	}
}
