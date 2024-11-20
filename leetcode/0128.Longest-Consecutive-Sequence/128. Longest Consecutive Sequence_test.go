package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0128 struct {
	nums []int
}

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		parameters0128
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0128{[]int{100, 4, 200, 1, 3, 2}},
			4,
		},
		{
			parameters0128{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			9,
		},
	}

	for _, test := range tests {
		t.Run("Test LongestConsecutiveSequence", func(t *testing.T) {
			// 完整輸入參數
			result := LongestConsecutiveSequence(test.parameters0128.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("LongestConsecutiveSequence(%+v) got %+v, want %+v", test.parameters0128, result, test.ans)
			}
		})
	}
}
