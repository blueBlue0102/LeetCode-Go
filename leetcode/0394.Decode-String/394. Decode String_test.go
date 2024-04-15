package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0394 struct {
	s string
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		parameters0394
		// 填入 function output type
		ans string
	}{
		// 填入 test case
		{
			parameters0394{"3[a]2[bc]"},
			"aaabcbc",
		},
		{
			parameters0394{"3[a2[c]]"},
			"accaccacc",
		},
		{
			parameters0394{"2[abc]3[cd]ef"},
			"abcabccdcdcdef",
		},
	}

	for _, test := range tests {
		t.Run("Test DecodeString", func(t *testing.T) {
			// 完整輸入參數
			result := DecodeString(test.parameters0394.s)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("DecodeString(%+v) got %+v, want %+v", test.parameters0394, result, test.ans)
			}
		})
	}
}
