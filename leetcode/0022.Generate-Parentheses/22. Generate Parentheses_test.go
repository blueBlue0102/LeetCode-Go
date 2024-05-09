package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 填入 function input type
type parameters0022 struct {
	n int
}

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		parameters0022
		// 填入 function output type
		ans []string
	}{
		// 填入 test case
		{
			parameters0022{1},
			[]string{"()"},
		},
		{
			parameters0022{3},
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, test := range tests {
		t.Run("Test GenerateParentheses", func(t *testing.T) {
			// 完整輸入參數
			result := GenerateParentheses(test.parameters0022.n)
			// compare 的方式需視情況調整
			sort.Strings(test.ans)
			sort.Strings(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("GenerateParentheses(%+v) got %+v, want %+v", test.parameters0022, result, test.ans)
			}
		})
	}
}
