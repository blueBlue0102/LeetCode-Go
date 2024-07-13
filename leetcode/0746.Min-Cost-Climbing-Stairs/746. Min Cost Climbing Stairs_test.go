package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0746 struct {
	cost []int
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		parameters0746
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0746{[]int{10, 15, 20}},
			15,
		},
		{
			parameters0746{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			6,
		},
	}

	for _, test := range tests {
		t.Run("Test MinCostClimbingStairs", func(t *testing.T) {
			// 完整輸入參數
			result := MinCostClimbingStairs(test.parameters0746.cost)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("MinCostClimbingStairs(%+v) got %+v, want %+v", test.parameters0746, result, test.ans)
			}
		})
	}
}
