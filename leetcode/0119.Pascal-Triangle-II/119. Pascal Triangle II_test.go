package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0119 struct {
	rowIndex int
}

func TestPascalTriangleII(t *testing.T) {
	tests := []struct {
		parameters0119
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0119{3},
			[]int{1, 3, 3, 1},
		},
		{
			parameters0119{0},
			[]int{1},
		},
		{
			parameters0119{1},
			[]int{1, 1},
		},
	}

	for _, test := range tests {
		t.Run("Test PascalTriangleII", func(t *testing.T) {
			// 完整輸入參數
			result := PascalTriangleII(test.parameters0119.rowIndex)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("PascalTriangleII(%+v) got %+v, want %+v", test.parameters0119, result, test.ans)
			}
		})
	}
}
