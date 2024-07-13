package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0198 struct {
	nums []int
}

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		parameters0198
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0198{[]int{1, 2, 3, 1}},
			4,
		},
		{
			parameters0198{[]int{2, 7, 9, 3, 1}},
			12,
		},
	}

	for _, test := range tests {
		t.Run("Test HouseRobber", func(t *testing.T) {
			// 完整輸入參數
			result := HouseRobber(test.parameters0198.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("HouseRobber(%+v) got %+v, want %+v", test.parameters0198, result, test.ans)
			}
		})
	}
}
