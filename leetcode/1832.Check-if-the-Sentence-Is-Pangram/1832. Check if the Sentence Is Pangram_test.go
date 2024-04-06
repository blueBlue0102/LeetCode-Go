package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters1832 struct {
	sentence string
}

func TestCheckiftheSentenceIsPangram(t *testing.T) {
	tests := []struct {
		parameters1832
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters1832{"thequickbrownfoxjumpsoverthelazydog"},
			true,
		},
		{
			parameters1832{"leetcode"},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test CheckiftheSentenceIsPangram", func(t *testing.T) {
			// 完整輸入參數
			result := CheckiftheSentenceIsPangram(test.parameters1832.sentence)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("CheckiftheSentenceIsPangram(%+v) got %+v, want %+v", test.parameters1832, result, test.ans)
			}
		})
	}
}
