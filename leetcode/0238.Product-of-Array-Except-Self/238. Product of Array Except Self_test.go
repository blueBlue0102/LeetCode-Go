package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0238 struct {
	nums []int
}

func TestProductofArrayExceptSelf(t *testing.T) {
	tests := []struct {
		parameters0238
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0238{[]int{1, 2, 3, 4}},
			[]int{24, 12, 8, 6},
		},
		{
			parameters0238{[]int{-1, 1, 0, -3, 3}},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run("Test ProductofArrayExceptSelf", func(t *testing.T) {
			// 完整輸入參數
			result := ProductofArrayExceptSelf(test.parameters0238.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ProductofArrayExceptSelf(%+v) got %+v, want %+v", test.parameters0238, result, test.ans)
			}
		})
	}
}
