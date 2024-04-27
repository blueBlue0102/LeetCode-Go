package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0912 struct {
	nums []int
}

func TestSortanArray(t *testing.T) {
	tests := []struct {
		parameters0912
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0912{[]int{5, 2, 3, 1}},
			[]int{1, 2, 3, 5},
		},
		{
			parameters0912{[]int{5, 1, 1, 2, 0, 0}},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, test := range tests {
		t.Run("Test SortanArray", func(t *testing.T) {
			// 完整輸入參數
			result := SortanArray(test.parameters0912.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SortanArray(%+v) got %+v, want %+v", test.parameters0912, result, test.ans)
			}
		})
	}
}
