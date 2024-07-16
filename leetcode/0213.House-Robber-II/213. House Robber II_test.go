package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0213 struct {
	nums []int
}

func TestHouseRobberII(t *testing.T) {
	tests := []struct {
		parameters0213
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0213{[]int{2, 3, 2}},
			3,
		},
		{
			parameters0213{[]int{1, 2, 3, 1}},
			4,
		},
		{
			parameters0213{[]int{1, 2, 3}},
			3,
		},
	}

	for _, test := range tests {
		t.Run("Test HouseRobberII", func(t *testing.T) {
			// 完整輸入參數
			result := HouseRobberII(test.parameters0213.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("HouseRobberII(%+v) got %+v, want %+v", test.parameters0213, result, test.ans)
			}
		})
	}
}
