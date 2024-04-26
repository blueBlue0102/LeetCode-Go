package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0069 struct {
	x int
}

func TestSqrtX(t *testing.T) {
	tests := []struct {
		parameters0069
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0069{4},
			2,
		},
		{
			parameters0069{8},
			2,
		},
		{
			parameters0069{9},
			3,
		},
		{
			parameters0069{0},
			0,
		},
		{
			parameters0069{2},
			1,
		},
		{
			parameters0069{1},
			1,
		},
	}

	for _, test := range tests {
		t.Run("Test SqrtX", func(t *testing.T) {
			// 完整輸入參數
			result := SqrtX(test.parameters0069.x)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SqrtX(%+v) got %+v, want %+v", test.parameters0069, result, test.ans)
			}
		})
	}
}
