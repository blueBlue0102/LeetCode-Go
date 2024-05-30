package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0033 struct {
	nums   []int
	target int
}

func TestSearchinRotatedSortedArray(t *testing.T) {
	tests := []struct {
		parameters0033
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0033{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			4,
		},
		{
			parameters0033{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			-1,
		},
		{
			parameters0033{[]int{1}, 0},
			-1,
		},
	}

	for _, test := range tests {
		t.Run("Test SearchinRotatedSortedArray", func(t *testing.T) {
			// 完整輸入參數
			result := SearchinRotatedSortedArray(test.parameters0033.nums, test.parameters0033.target)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SearchinRotatedSortedArray(%+v) got %+v, want %+v", test.parameters0033, result, test.ans)
			}
		})
	}
}
