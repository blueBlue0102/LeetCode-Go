package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters1047 struct {
	s string
}

func TestRemoveAllAdjacentDuplicatesInString(t *testing.T) {
	tests := []struct {
		parameters1047
		// 填入 function output type
		ans string
	}{
		// 填入 test case
		{
			parameters1047{"abbaca"},
			"ca",
		},
		{
			parameters1047{"azxxzy"},
			"ay",
		},
		{
			parameters1047{"a"},
			"a",
		},
		{
			parameters1047{"aaa"},
			"a",
		},
		{
			parameters1047{"aaaaaaaaaa"},
			"",
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveAllAdjacentDuplicatesInString", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveAllAdjacentDuplicatesInString(test.parameters1047.s)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("RemoveAllAdjacentDuplicatesInString(%+v) got %+v, want %+v", test.parameters1047, result, test.ans)
			}
		})
	}
}
