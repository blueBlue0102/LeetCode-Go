package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0704 struct {
	nums   []int
	target int
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		parameters0704
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0704{[]int{-1, 0, 3, 5, 9, 12}, 9},
			4,
		},
		{
			parameters0704{[]int{-1, 0, 3, 5, 9, 12}, 2},
			-1,
		},
	}

	for _, test := range tests {
		t.Run("Test BinarySearch", func(t *testing.T) {
			// 完整輸入參數
			result := BinarySearch(test.parameters0704.nums, test.parameters0704.target)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("BinarySearch(%+v) got %+v, want %+v", test.parameters0704, result, test.ans)
			}
		})
	}
}
