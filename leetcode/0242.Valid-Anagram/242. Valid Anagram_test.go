package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0242 struct {
	s string
	t string
}

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		parameters0242
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0242{"anagram", "nagaram"},
			true,
		},
		{
			parameters0242{"rat", "car"},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test ValidAnagram", func(t *testing.T) {
			// 完整輸入參數
			result := ValidAnagram(test.parameters0242.s, test.parameters0242.t)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ValidAnagram(%+v) got %+v, want %+v", test.parameters0242, result, test.ans)
			}
		})
	}
}
