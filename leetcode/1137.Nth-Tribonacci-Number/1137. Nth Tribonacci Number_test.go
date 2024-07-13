package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters1137 struct {
	n int
}

func TestNthTribonacciNumber(t *testing.T) {
	tests := []struct {
		parameters1137
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters1137{4},
			4,
		},
		{
			parameters1137{25},
			1389537,
		},
	}

	for _, test := range tests {
		t.Run("Test NthTribonacciNumber", func(t *testing.T) {
			// 完整輸入參數
			result := NthTribonacciNumber(test.parameters1137.n)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("NthTribonacciNumber(%+v) got %+v, want %+v", test.parameters1137, result, test.ans)
			}
		})
	}
}
