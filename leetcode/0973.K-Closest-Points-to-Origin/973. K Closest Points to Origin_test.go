package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0973 struct {
	points [][]int
	k      int
}

func TestKClosestPointstoOrigin(t *testing.T) {
	tests := []struct {
		parameters0973
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0973{[][]int{{1, 3}, {-2, 2}}, 1},
			utils.Sort2DArray([][]int{{-2, 2}}),
		},
		{
			parameters0973{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
			utils.Sort2DArray([][]int{{3, 3}, {-2, 4}}),
		},
		{
			parameters0973{[][]int{{0, 1}, {1, 0}}, 2},
			utils.Sort2DArray([][]int{{0, 1}, {1, 0}}),
		},
	}

	for _, test := range tests {
		t.Run("Test KClosestPointstoOrigin", func(t *testing.T) {
			// 完整輸入參數
			result := utils.Sort2DArray(KClosestPointstoOrigin(test.parameters0973.points, test.parameters0973.k))
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("KClosestPointstoOrigin(%+v) got %+v, want %+v", test.parameters0973, result, test.ans)
			}
		})
	}
}
