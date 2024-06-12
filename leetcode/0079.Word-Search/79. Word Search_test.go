package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0079 struct {
	board [][]byte
	word  string
}

func TestWordSearch(t *testing.T) {
	tests := []struct {
		parameters0079
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0079{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"},
			true,
		},
		{
			parameters0079{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"},
			true,
		},
		{
			parameters0079{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test WordSearch", func(t *testing.T) {
			// 完整輸入參數
			result := WordSearch(test.parameters0079.board, test.parameters0079.word)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("WordSearch(%+v) got %+v, want %+v", test.parameters0079, result, test.ans)
			}
		})
	}
}
