package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0202 struct {
	n int
}

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		parameters0202
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0202{19},
			true,
		},
		{
			parameters0202{2},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test HappyNumber", func(t *testing.T) {
			// 完整輸入參數
			result := HappyNumber(test.parameters0202.n)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("HappyNumber(%+v) got %+v, want %+v", test.parameters0202, result, test.ans)
			}
		})
	}
}
