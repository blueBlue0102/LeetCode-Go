package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0323 struct {
	n     int
	edges [][]int
}

func TestNumberofConnectedComponentsinanUndirectedGraph(t *testing.T) {
	tests := []struct {
		parameters0323
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0323{5, [][]int{{0, 1}, {1, 2}, {3, 4}}},
			2,
		},
		{
			parameters0323{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
			1,
		},
	}

	for _, test := range tests {
		t.Run("Test NumberofConnectedComponentsinanUndirectedGraph", func(t *testing.T) {
			// 完整輸入參數
			result := NumberofConnectedComponentsinanUndirectedGraph(test.parameters0323.n, test.parameters0323.edges)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("NumberofConnectedComponentsinanUndirectedGraph(%+v) got %+v, want %+v", test.parameters0323, result, test.ans)
			}
		})
	}
}
