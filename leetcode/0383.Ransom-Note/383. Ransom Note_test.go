package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0383 struct {
	ransomNote string
	magazine   string
}

func TestRansomNote(t *testing.T) {
	tests := []struct {
		parameters0383
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0383{"a", "b"},
			false,
		},
		{
			parameters0383{"aa", "ab"},
			false,
		},
		{
			parameters0383{"aa", "aab"},
			true,
		},
	}

	for _, test := range tests {
		t.Run("Test RansomNote", func(t *testing.T) {
			// 完整輸入參數
			result := RansomNote(test.parameters0383.ransomNote, test.parameters0383.magazine)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("RansomNote(%+v) got %+v, want %+v", test.parameters0383, result, test.ans)
			}
		})
	}
}
