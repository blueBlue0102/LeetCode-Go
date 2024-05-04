package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters1079 struct {
	tiles string
}

func TestLetterTilePossibilities(t *testing.T) {
	tests := []struct {
		parameters1079
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters1079{"AAB"},
			8,
		},
		{
			parameters1079{"AAABBC"},
			188,
		},
		{
			parameters1079{"V"},
			1,
		},
	}

	for _, test := range tests {
		t.Run("Test LetterTilePossibilities", func(t *testing.T) {
			// 完整輸入參數
			result := LetterTilePossibilities(test.parameters1079.tiles)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LetterTilePossibilities(%+v) got %+v, want %+v", test.parameters1079, result, test.ans)
			}
		})
	}
}
