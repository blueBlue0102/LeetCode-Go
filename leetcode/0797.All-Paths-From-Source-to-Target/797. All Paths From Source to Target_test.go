package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0797 struct {
	graph [][]int
}

func TestAllPathsFromSourcetoTarget(t *testing.T) {
	tests := []struct {
		parameters0797
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0797{[][]int{{1, 2}, {3}, {3}, {}}},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			parameters0797{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			parameters0797{[][]int{{2}, {2}, {}}},
			[][]int{{0, 2}},
		},
		{
			parameters0797{[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}},
			[][]int{{0, 3, 6, 7}, {0, 3, 4, 7}, {0, 3, 4, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 7}, {0, 1, 4, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 6, 7}, {0, 1, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}, {0, 1, 5, 6, 7}},
		},
	}

	for _, test := range tests {
		t.Run("Test AllPathsFromSourcetoTarget", func(t *testing.T) {
			test.ans = utils.Sort2DArray(test.ans)
			// 完整輸入參數
			result := utils.Sort2DArray(AllPathsFromSourcetoTarget(test.parameters0797.graph))
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("AllPathsFromSourcetoTarget(%+v) got %+v, want %+v", test.parameters0797, result, test.ans)
			}
		})
	}
}
