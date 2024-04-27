package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters2545 struct {
	score [][]int
	k     int
}

func TestSorttheStudentsbyTheirKthScore(t *testing.T) {
	tests := []struct {
		parameters2545
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters2545{[][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2},
			[][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}},
		},
		{
			parameters2545{[][]int{{3, 4}, {5, 6}}, 0},
			[][]int{{5, 6}, {3, 4}},
		},
	}

	for _, test := range tests {
		t.Run("Test SorttheStudentsbyTheirKthScore", func(t *testing.T) {
			// 完整輸入參數
			result := SorttheStudentsbyTheirKthScore(test.parameters2545.score, test.parameters2545.k)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("SorttheStudentsbyTheirKthScore(%+v) got %+v, want %+v", test.parameters2545, result, test.ans)
			}
		})
	}
}
