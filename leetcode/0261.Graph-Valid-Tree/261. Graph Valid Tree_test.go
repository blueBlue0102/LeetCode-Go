package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0261 struct {
	n     int
	edges [][]int
}

func TestGraphValidTree(t *testing.T) {
	tests := []struct {
		parameters0261
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0261{5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}},
			true,
		},
		{
			parameters0261{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}},
			false,
		},
		{
			parameters0261{4, [][]int{{0, 1}, {2, 3}}},
			false,
		},
	}

	for _, test := range tests {
		t.Run("Test GraphValidTree", func(t *testing.T) {
			// 完整輸入參數
			result := GraphValidTree(test.parameters0261.n, test.parameters0261.edges)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("GraphValidTree(%+v) got %+v, want %+v", test.parameters0261, result, test.ans)
			}
		})
	}
}
