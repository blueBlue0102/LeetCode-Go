package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0417 struct {
	heights [][]int
}

func TestPacificAtlanticWaterFlow(t *testing.T) {
	tests := []struct {
		parameters0417
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0417{[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			parameters0417{[][]int{{1}}},
			[][]int{{0, 0}},
		},
	}

	for _, test := range tests {
		t.Run("Test PacificAtlanticWaterFlow", func(t *testing.T) {
			// 完整輸入參數
			result := PacificAtlanticWaterFlow(test.parameters0417.heights)
			// compare 的方式需視情況調整
			utils.Sort2DArray(test.ans)
			utils.Sort2DArray(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("PacificAtlanticWaterFlow(%+v) got %+v, want %+v", test.parameters0417, result, test.ans)
			}
		})
	}
}
