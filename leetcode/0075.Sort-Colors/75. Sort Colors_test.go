package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0075 struct {
	nums []int
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		parameters0075
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0075{[]int{2, 0, 2, 1, 1, 0}},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			parameters0075{[]int{2, 0, 1}},
			[]int{0, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run("Test SortColors", func(t *testing.T) {
			// 完整輸入參數
			SortColors(test.parameters0075.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(test.parameters0075.nums, test.ans) {
				t.Errorf("SortColors(%+v) got %+v, want %+v", test.parameters0075, test.parameters0075.nums, test.ans)
			}
		})
	}
}
