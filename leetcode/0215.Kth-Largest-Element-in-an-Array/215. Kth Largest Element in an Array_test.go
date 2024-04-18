package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0215 struct {
	nums []int
	k    int
}

func TestKthLargestElementinanArray(t *testing.T) {
	tests := []struct {
		parameters0215
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0215{[]int{3, 2, 1, 5, 6, 4}, 2},
			5,
		},
		{
			parameters0215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			4,
		},
		{
			parameters0215{[]int{2, 1}, 1},
			2,
		},
		{
			parameters0215{[]int{5, 2, 4, 1, 3, 6, 0}, 4},
			3,
		},
	}

	for _, test := range tests {
		t.Run("Test KthLargestElementinanArray", func(t *testing.T) {
			// 完整輸入參數
			result := KthLargestElementinanArray(test.parameters0215.nums, test.parameters0215.k)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("KthLargestElementinanArray(%+v) got %+v, want %+v", test.parameters0215, result, test.ans)
			}
		})
	}
}
