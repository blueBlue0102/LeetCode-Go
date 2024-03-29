package leetcode

import (
	"reflect"
	"slices"
	"testing"
)

// 填入 function input type
type parameters1002 struct {
	words []string
}

func TestFindCommonCharacters(t *testing.T) {
	tests := []struct {
		parameters1002
		// 填入 function output type
		ans []string
	}{
		// 填入 test case
		{
			parameters1002{[]string{"bella", "label", "roller"}},
			[]string{"e", "l", "l"},
		},
		{
			parameters1002{[]string{"cool", "lock", "cook"}},
			[]string{"c", "o"},
		},
		{
			parameters1002{[]string{"aaa", "bbb", "ccc"}},
			[]string{},
		},
		{
			parameters1002{[]string{"我不要", "我要怎麼了", "我想要"}},
			[]string{"我", "要"},
		},
	}

	for _, test := range tests {
		t.Run("Test FindCommonCharacters", func(t *testing.T) {
			// 完整輸入參數
			result := FindCommonCharacters(test.parameters1002.words)
			slices.Sort(result)
			slices.Sort(test.ans)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("FindCommonCharacters(%+v) got %+v, want %+v", test.parameters1002, result, test.ans)
			}
		})
	}
}
