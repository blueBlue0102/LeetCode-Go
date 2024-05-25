package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0733 struct {
	image [][]int
	sr    int
	sc    int
	color int
}

func TestFloodFill(t *testing.T) {
	tests := []struct {
		parameters0733
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0733{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2},
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			parameters0733{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0},
			[][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, test := range tests {
		t.Run("Test FloodFill", func(t *testing.T) {
			// 完整輸入參數
			result := FloodFill(test.parameters0733.image, test.parameters0733.sr, test.parameters0733.sc, test.parameters0733.color)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("FloodFill(%+v) got %+v, want %+v", test.parameters0733, result, test.ans)
			}
		})
	}
}
