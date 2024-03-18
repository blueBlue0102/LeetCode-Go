package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0059 struct {
	n int
}

func TestSpiralMatrixII(t *testing.T) {
	tests := []struct {
		parameters0059
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0059{n: 3},
			[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			parameters0059{n: 1},
			[][]int{{1}},
		},
		{
			parameters0059{n: 2},
			[][]int{{1, 2}, {4, 3}},
		},
		{
			parameters0059{n: 4},
			[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		},
	}

	for _, test := range tests {
		t.Run("Test SpiralMatrixII", func(t *testing.T) {
			// 完整輸入參數
			result := SpiralMatrixII(test.parameters0059.n)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SpiralMatrixII(%+v) got %+v, want %+v", test.parameters0059, result, test.ans)
			}
		})
	}
}
