package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0020 struct {
	s string
}

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		parameters0020
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0020{"()"},
			true,
		},
		{
			parameters0020{"()[]{}"},
			true,
		},
		{
			parameters0020{"(]"},
			false,
		},
		{
			parameters0020{"[(])"},
			false,
		},
		{
			parameters0020{"((())))"},
			false,
		},
		{
			parameters0020{")"},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test ValidParentheses", func(t *testing.T) {
			// 完整輸入參數
			result := ValidParentheses(test.parameters0020.s)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ValidParentheses(%+v) got %+v, want %+v", test.parameters0020, result, test.ans)
			}
		})
	}
}
