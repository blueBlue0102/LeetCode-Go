package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters2225 struct {
	matches [][]int
}

func TestFindPlayersWithZeroorOneLosses(t *testing.T) {
	tests := []struct {
		parameters2225
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters2225{[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}},
			[][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			parameters2225{[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}},
			[][]int{{1, 2, 5, 6}, {}},
		},
	}

	for _, test := range tests {
		t.Run("Test FindPlayersWithZeroorOneLosses", func(t *testing.T) {
			// 完整輸入參數
			result := FindPlayersWithZeroorOneLosses(test.parameters2225.matches)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("FindPlayersWithZeroorOneLosses(%+v) got %+v, want %+v", test.parameters2225, result, test.ans)
			}
		})
	}
}
