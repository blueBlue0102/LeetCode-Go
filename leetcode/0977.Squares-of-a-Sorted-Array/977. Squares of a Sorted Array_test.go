package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0977 struct {
	nums []int
}

func TestSquaresofaSortedArray(t *testing.T) {
	tests := []struct {
		parameters0977
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0977{nums: []int{-4, -1, 0, 3, 10}},
			[]int{0, 1, 9, 16, 100},
		},
		{
			parameters0977{nums: []int{-7, -3, 2, 3, 11}},
			[]int{4, 9, 9, 49, 121},
		},
	}

	for _, test := range tests {
		t.Run("Test SquaresofaSortedArray", func(t *testing.T) {
			// 完整輸入參數
			result := SquaresofaSortedArray(test.parameters0977.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SquaresofaSortedArray(%+v) got %+v, want %+v", test.parameters0977, result, test.ans)
			}
		})
	}
}
