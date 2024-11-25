package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0015 struct {
	nums []int
}

func Test3Sum(t *testing.T) {
	tests := []struct {
		parameters0015
		// 填入 function output type
		ans [][]int
	}{
		// 填入 test case
		{
			parameters0015{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			parameters0015{[]int{0, 1, 1}},
			[][]int{},
		},
		{
			parameters0015{[]int{0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
		{
			parameters0015{[]int{0, 0, 0, 0}},
			[][]int{{0, 0, 0}},
		},
	}

	for _, test := range tests {
		t.Run("Test 3Sum", func(t *testing.T) {
			// 完整輸入參數
			result := ThreeSum(test.parameters0015.nums)
			// compare 的方式需視情況調整
			utils.Sort2DArray(result)
			utils.Sort2DArray(test.ans)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("3Sum(%+v) got %+v, want %+v", test.parameters0015, result, test.ans)
			}
		})
	}
}
