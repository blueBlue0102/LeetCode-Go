package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0121 struct {
	prices []int
}

func TestBestTimetoBuyandSellStock(t *testing.T) {
	tests := []struct {
		parameters0121
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0121{[]int{7, 1, 5, 3, 6, 4}},
			5,
		},
		{
			parameters0121{[]int{7, 6, 4, 3, 1}},
			0,
		},
	}

	for _, test := range tests {
		t.Run("Test BestTimetoBuyandSellStock", func(t *testing.T) {
			// 完整輸入參數
			result := BestTimetoBuyandSellStock(test.parameters0121.prices)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BestTimetoBuyandSellStock(%+v) got %+v, want %+v", test.parameters0121, result, test.ans)
			}
		})
	}
}
