package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0001 struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		parameters0001
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0001{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			parameters0001{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			parameters0001{[]int{3, 3}, 6},
			[]int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run("Test TwoSum", func(t *testing.T) {
			// 完整輸入參數
			result := TwoSum(test.parameters0001.nums, test.parameters0001.target)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("TwoSum(%+v) got %+v, want %+v", test.parameters0001, result, test.ans)
			}
		})
	}
}
