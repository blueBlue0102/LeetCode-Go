package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0003 struct {
	s string
}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		parameters0003
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0003{"abcabcbb"},
			3,
		},
		{
			parameters0003{"bbbbb"},
			1,
		},
		{
			parameters0003{"pwwkew"},
			3,
		},
	}

	for _, test := range tests {
		t.Run("Test LongestSubstringWithoutRepeatingCharacters", func(t *testing.T) {
			// 完整輸入參數
			result := LongestSubstringWithoutRepeatingCharacters(test.parameters0003.s)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LongestSubstringWithoutRepeatingCharacters(%+v) got %+v, want %+v", test.parameters0003, result, test.ans)
			}
		})
	}
}
