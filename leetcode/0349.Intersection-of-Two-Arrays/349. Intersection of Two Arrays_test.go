package leetcode

import (
	"reflect"
	"slices"
	"testing"
)

// 填入 function input type
type parameters0349 struct {
	nums1 []int
	nums2 []int
}

func TestIntersectionofTwoArrays(t *testing.T) {
	tests := []struct {
		parameters0349
		// 填入 function output type
		ans []int
	}{
		// 填入 test case
		{
			parameters0349{[]int{1, 2, 2, 1}, []int{2, 2}},
			[]int{2},
		},
		{
			parameters0349{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			[]int{4, 9},
		},
	}

	for _, test := range tests {
		t.Run("Test IntersectionofTwoArrays", func(t *testing.T) {
			// 完整輸入參數
			result := IntersectionofTwoArrays(test.parameters0349.nums1, test.parameters0349.nums2)
			// compare 的方式需視情況調整
			slices.Sort(result)
			slices.Sort(test.ans)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("IntersectionofTwoArrays(%+v) got %+v, want %+v", test.parameters0349, result, test.ans)
			}
		})
	}
}
