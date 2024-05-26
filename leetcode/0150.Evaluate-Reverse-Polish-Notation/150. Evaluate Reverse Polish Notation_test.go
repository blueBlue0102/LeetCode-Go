package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0150 struct {
	tokens []string
}

func TestEvaluateReversePolishNotation(t *testing.T) {
	tests := []struct {
		parameters0150
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0150{[]string{"2", "1", "+", "3", "*"}},
			9,
		},
		{
			parameters0150{[]string{"4", "13", "5", "/", "+"}},
			6,
		},
		{
			parameters0150{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			22,
		},
	}

	for _, test := range tests {
		t.Run("Test EvaluateReversePolishNotation", func(t *testing.T) {
			// 完整輸入參數
			result := EvaluateReversePolishNotation(test.parameters0150.tokens)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("EvaluateReversePolishNotation(%+v) got %+v, want %+v", test.parameters0150, result, test.ans)
			}
		})
	}
}
