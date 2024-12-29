package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0424 struct {
	s string
	k int
}

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	tests := []struct {
		parameters0424
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0424{"ABAB", 2},
			4,
		},
		{
			parameters0424{"AABABBA", 1},
			4,
		},
		{
			parameters0424{"ABBB", 2},
			4,
		},
		{
			parameters0424{"AAAACBBBBAAAAAA", 5},
			15,
		},
	}

	for _, test := range tests {
		t.Run("Test LongestRepeatingCharacterReplacement", func(t *testing.T) {
			// 完整輸入參數
			result := LongestRepeatingCharacterReplacement(test.parameters0424.s, test.parameters0424.k)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LongestRepeatingCharacterReplacement(%+v) got %+v, want %+v", test.parameters0424, result, test.ans)
			}
		})
	}
}
