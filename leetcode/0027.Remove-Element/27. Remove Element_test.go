package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0027 struct {
	nums []int
	val  int
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		parameters0027
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0027{[]int{3, 2, 2, 3}, 3},
			2,
		},
		{
			parameters0027{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			5,
		},
	}

	for _, test := range tests {
		t.Run("Test RemoveElement", func(t *testing.T) {
			// 完整輸入參數
			result := RemoveElement(test.parameters0027.nums, test.parameters0027.val)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("RemoveElement(%+v) got %+v, want %+v", test.parameters0027, result, test.ans)
			}
		})
	}
}
