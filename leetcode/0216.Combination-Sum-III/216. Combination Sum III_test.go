package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0216 struct {
	k int
	n int
}

func TestCombinationSumIII(t *testing.T) {
	tests := []struct {
		parameters0216
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0216{3, 7},
			[][]int{{1, 2, 4}},
		},
		{
			parameters0216{3, 9},
			[][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			parameters0216{4, 1},
			[][]int{},
		},
		{
			parameters0216{4, 24},
			[][]int{{1, 6, 8, 9}, {2, 5, 8, 9}, {2, 6, 7, 9}, {3, 4, 8, 9}, {3, 5, 7, 9}, {3, 6, 7, 8}, {4, 5, 6, 9}, {4, 5, 7, 8}},
		},
	}

	for _, test := range tests {
		t.Run("Test CombinationSumIII", func(t *testing.T) {
			// 完整輸入參數
			result := CombinationSumIII(test.parameters0216.k, test.parameters0216.n)
			// compare 的方式需視情況調整
			test.ans = utils.Sort2DIntArray(test.ans)
			result = utils.Sort2DIntArray(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("CombinationSumIII(%+v) got %+v, want %+v", test.parameters0216, result, test.ans)
			}
		})
	}
}
