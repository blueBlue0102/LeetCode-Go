package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0153 struct {
	nums []int
}

func TestFindMinimuminRotatedSortedArray(t *testing.T) {
	tests := []struct {
		parameters0153
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0153{[]int{3, 4, 5, 1, 2}},
			1,
		},
		{
			parameters0153{[]int{4, 5, 6, 7, 0, 1, 2}},
			0,
		},
		{
			parameters0153{[]int{11, 13, 15, 17}},
			11,
		},
		{
			parameters0153{[]int{2, 1}},
			1,
		},
	}

	for _, test := range tests {
		t.Run("Test FindMinimuminRotatedSortedArray", func(t *testing.T) {
			// 完整輸入參數
			result := FindMinimuminRotatedSortedArray(test.parameters0153.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("FindMinimuminRotatedSortedArray(%+v) got %+v, want %+v", test.parameters0153, result, test.ans)
			}
		})
	}
}
