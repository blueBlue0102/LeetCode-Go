package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 填入 function input type
type parameters0017 struct {
	digits string
}

func TestLetterCombinationsofaPhoneNumber(t *testing.T) {
	tests := []struct {
		parameters0017
		// 填入 function output type
		ans []string
	}{
		// 填入 test case
		{
			parameters0017{"23"},
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			parameters0017{""},
			[]string{},
		},
		{
			parameters0017{"2"},
			[]string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run("Test LetterCombinationsofaPhoneNumber", func(t *testing.T) {
			sort.Strings(test.ans)
			// 完整輸入參數
			result := LetterCombinationsofaPhoneNumber(test.parameters0017.digits)
			sort.Strings(result)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LetterCombinationsofaPhoneNumber(%+v) got %+v, want %+v", test.parameters0017, result, test.ans)
			}
		})
	}
}
