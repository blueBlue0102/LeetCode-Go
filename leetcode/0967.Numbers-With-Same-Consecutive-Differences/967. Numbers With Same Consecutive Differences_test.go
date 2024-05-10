package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 填入 function input type
type parameters0967 struct {
	n int
	k int
}

func TestNumbersWithSameConsecutiveDifferences(t *testing.T) {
	tests := []struct {
		parameters0967
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0967{3, 7},
			[]int{181, 292, 707, 818, 929},
		},
		{
			parameters0967{2, 1},
			[]int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98},
		},
		{
			parameters0967{2, 0},
			[]int{11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
		{
			parameters0967{3, 1},
			[]int{101, 121, 123, 210, 212, 232, 234, 321, 323, 343, 345, 432, 434, 454, 456, 543, 545, 565, 567, 654, 656, 676, 678, 765, 767, 787, 789, 876, 878, 898, 987, 989},
		},
	}

	for _, test := range tests {
		t.Run("Test NumbersWithSameConsecutiveDifferences", func(t *testing.T) {
			// 完整輸入參數
			result := NumbersWithSameConsecutiveDifferences(test.parameters0967.n, test.parameters0967.k)
			// compare 的方式需視情況調整
			sort.Ints(test.ans)
			sort.Ints(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("NumbersWithSameConsecutiveDifferences(%+v) got %+v, want %+v", test.parameters0967, result, test.ans)
			}
		})
	}
}
