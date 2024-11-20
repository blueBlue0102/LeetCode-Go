package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0005 struct {
	s string
}

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		parameters0005
		// 填入 function output type
		ans string
	}{
		// 填入 test case
		{
			parameters0005{"babad"},
			"bab",
		},
		{
			parameters0005{"cbbd"},
			"bb",
		},
		{
			parameters0005{"abcd"},
			"a",
		},
	}

	for _, test := range tests {
		t.Run("Test LongestPalindromicSubstring", func(t *testing.T) {
			// 完整輸入參數
			result := LongestPalindromicSubstring(test.parameters0005.s)
			// compare 的方式需視情況調整
			// 由於答案可能不只一組，所以用長度來辨別
			if !reflect.DeepEqual(len(result), len(test.ans)) {
				t.Errorf("LongestPalindromicSubstring(%+v) got %+v, want %+v", test.parameters0005, result, test.ans)
			}
		})
	}
}
