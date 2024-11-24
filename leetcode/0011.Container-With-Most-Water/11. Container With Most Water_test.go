package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0011 struct {
	height []int
}

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		parameters0011
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0011{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			49,
		},
		{
			parameters0011{[]int{1, 1}},
			1,
		},
	}

	for _, test := range tests {
		t.Run("Test ContainerWithMostWater", func(t *testing.T) {
			// 完整輸入參數
			result := ContainerWithMostWater(test.parameters0011.height)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("ContainerWithMostWater(%+v) got %+v, want %+v", test.parameters0011, result, test.ans)
			}
		})
	}
}
