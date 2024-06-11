package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0039 struct {
	candidates []int
	target     int
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		parameters0039
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0039{[]int{2, 3, 6, 7}, 7},
			[][]int{{2, 2, 3}, {7}},
		},
		{
			parameters0039{[]int{2, 3, 5}, 8},
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			parameters0039{[]int{2}, 1},
			[][]int{},
		},
		{
			parameters0039{[]int{8, 7, 4, 3}, 11},
			[][]int{{8, 3}, {7, 4}, {4, 4, 3}},
		},
	}

	for _, test := range tests {
		t.Run("Test CombinationSum", func(t *testing.T) {
			// 完整輸入參數
			result := CombinationSum(test.parameters0039.candidates, test.parameters0039.target)
			// compare 的方式需視情況調整
			result = utils.Sort2DArray(result)
			test.ans = utils.Sort2DArray(test.ans)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("CombinationSum(%+v) got %+v, want %+v", test.parameters0039, result, test.ans)
			}
		})
	}
}
