package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0217 struct {
	nums []int
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		parameters0217
		// 填入 function output type
		ans bool
	}{
		// 填入 test case
		{
			parameters0217{[]int{1, 2, 3, 1}},
			true,
		},
		{
			parameters0217{[]int{1, 2, 3, 4}},
			false,
		},
		{
			parameters0217{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			true,
		},
	}

	for _, test := range tests {
		t.Run("Test ContainsDuplicate", func(t *testing.T) {
			// 完整輸入參數
			result := ContainsDuplicate(test.parameters0217.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ContainsDuplicate(%+v) got %+v, want %+v", test.parameters0217, result, test.ans)
			}
		})
	}
}
