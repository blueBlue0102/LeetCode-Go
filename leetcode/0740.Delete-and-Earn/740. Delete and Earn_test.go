package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0740 struct {
	nums []int
}

func TestDeleteandEarn(t *testing.T) {
	tests := []struct {
		parameters0740
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0740{[]int{3, 4, 2}},
			6,
		},
		{
			parameters0740{[]int{2, 2, 3, 3, 3, 4}},
			9,
		},
	}

	for _, test := range tests {
		t.Run("Test DeleteandEarn", func(t *testing.T) {
			// 完整輸入參數
			result := DeleteandEarn(test.parameters0740.nums)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("DeleteandEarn(%+v) got %+v, want %+v", test.parameters0740, result, test.ans)
			}
		})
	}
}
