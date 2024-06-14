package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0200 struct {
	grid [][]byte
}

func TestNumberofIslands(t *testing.T) {
	tests := []struct {
		parameters0200
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0200{[][]byte{{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'}}},
			1,
		},
		{
			parameters0200{[][]byte{{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'}}},
			3,
		},
	}

	for _, test := range tests {
		t.Run("Test NumberofIslands", func(t *testing.T) {
			// 完整輸入參數
			result := NumberofIslands(test.parameters0200.grid)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("NumberofIslands(%+v) got %+v, want %+v", test.parameters0200, result, test.ans)
			}
		})
	}
}
