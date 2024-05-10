package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0051 struct {
	n int
}

func TestNQueens(t *testing.T) {
	tests := []struct {
		parameters0051
		// 填入 function output type
		ans [][]string
	}{
		// 填入 test case
		{
			parameters0051{4},
			[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			parameters0051{1},
			[][]string{{"Q"}},
		},
	}

	for _, test := range tests {
		t.Run("Test NQueens", func(t *testing.T) {
			// 完整輸入參數
			result := NQueens(test.parameters0051.n)
			// compare 的方式需視情況調整
			utils.Sort2DArray(test.ans)
			utils.Sort2DArray(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("NQueens(%+v) got %+v, want %+v", test.parameters0051, result, test.ans)
			}
		})
	}
}
