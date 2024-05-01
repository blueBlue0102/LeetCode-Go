package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0046 struct {
	nums []int
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		parameters0046
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0046{[]int{1, 2, 3}},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			parameters0046{[]int{0, 1}},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			parameters0046{[]int{1}},
			[][]int{{1}},
		},
	}

	for _, test := range tests {
		t.Run("Test Permutations", func(t *testing.T) {
			test.ans = utils.Sort2DIntArray(test.ans)
			// 完整輸入參數
			result := utils.Sort2DIntArray(Permutations(test.parameters0046.nums))
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("Permutations(%+v) got %+v, want %+v", test.parameters0046, result, test.ans)
			}
		})
	}
}
