package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters1636 struct {
	nums []int
}

func TestSortArraybyIncreasingFrequency(t *testing.T) {
	tests := []struct {
		parameters1636
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters1636{[]int{1, 1, 2, 2, 2, 3}},
			[]int{3, 1, 1, 2, 2, 2},
		},
		{
			parameters1636{[]int{2, 3, 1, 3, 2}},
			[]int{1, 3, 3, 2, 2},
		},
		{
			parameters1636{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
			[]int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run("Test SortArraybyIncreasingFrequency", func(t *testing.T) {
			// 完整輸入參數
			result := SortArraybyIncreasingFrequency(test.parameters1636.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SortArraybyIncreasingFrequency(%+v) got %+v, want %+v", test.parameters1636, result, test.ans)
			}
		})
	}
}
