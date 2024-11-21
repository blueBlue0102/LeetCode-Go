package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 填入 function input type
type parameters0347 struct {
	nums []int
	k    int
}

func TestTopKFrequentElements(t *testing.T) {
	tests := []struct {
		parameters0347
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0347{[]int{1, 1, 1, 2, 2, 3}, 2},
			[]int{1, 2},
		},
		{
			parameters0347{[]int{1}, 1},
			[]int{1},
		},
		{
			parameters0347{[]int{6, 6, 6, 0, 6, 6, 6}, 1},
			[]int{6},
		},
	}

	for _, test := range tests {
		t.Run("Test TopKFrequentElements", func(t *testing.T) {
			// 完整輸入參數
			result := TopKFrequentElements(test.parameters0347.nums, test.parameters0347.k)
			// compare 的方式需視情況調整
			sort.Ints(result)
			sort.Ints(test.ans)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("TopKFrequentElements(%+v) got %+v, want %+v", test.parameters0347, result, test.ans)
			}
		})
	}
}
