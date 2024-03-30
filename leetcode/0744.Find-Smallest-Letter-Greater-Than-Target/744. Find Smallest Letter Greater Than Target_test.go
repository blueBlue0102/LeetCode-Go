package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0744 struct {
	letters []byte
	target  byte
}

func TestFindSmallestLetterGreaterThanTarget(t *testing.T) {
	tests := []struct {
		parameters0744
		// 填入 function output type
		ans byte
	}{
		// 填入 test case
		{
			parameters0744{[]byte{'c', 'f', 'j'}, 'a'},
			'c',
		},
		{
			parameters0744{[]byte{'c', 'f', 'j'}, 'c'},
			'f',
		},
		{
			parameters0744{[]byte{'x', 'x', 'y', 'y'}, 'z'},
			'x',
		},
	}

	for _, test := range tests {
		t.Run("Test FindSmallestLetterGreaterThanTarget", func(t *testing.T) {
			// 完整輸入參數
			result := FindSmallestLetterGreaterThanTarget(test.parameters0744.letters, test.parameters0744.target)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("FindSmallestLetterGreaterThanTarget(%+v) got %+v, want %+v", test.parameters0744, result, test.ans)
			}
		})
	}
}
