package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0070 struct {
	n int
}

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		parameters0070
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0070{2},
			2,
		},
		{
			parameters0070{3},
			3,
		},
		{
			parameters0070{4},
			5,
		},
		{
			parameters0070{5},
			8,
		},
		{
			parameters0070{39},
			102334155,
		},
	}

	for _, test := range tests {
		t.Run("Test ClimbingStairs", func(t *testing.T) {
			// 完整輸入參數
			result := ClimbingStairs(test.parameters0070.n)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ClimbingStairs(%+v) got %+v, want %+v", test.parameters0070, result, test.ans)
			}
		})
	}
}
