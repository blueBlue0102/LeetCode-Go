package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0125 struct {
	s string
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		parameters0125
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0125{"A man, a plan, a canal: Panama"},
			true,
		},
		{
			parameters0125{"race a car"},
			false,
		},
		{
			parameters0125{" "},
			true,
		},
	}

	for _, test := range tests {
		t.Run("Test ValidPalindrome", func(t *testing.T) {
			// 完整輸入參數
			result := ValidPalindrome(test.parameters0125.s)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ValidPalindrome(%+v) got %+v, want %+v", test.parameters0125, result, test.ans)
			}
		})
	}
}
