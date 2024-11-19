package leetcode

import (
	"reflect"
	"testing"

	"github.com/blueBlue0102/LeetCode-Go/internal/utils"
)

// 填入 function input type
type parameters0049 struct {
	strs []string
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		parameters0049
		// 填入 function output type
		ans [][]string
	}{
		// 填入 test case
		{
			parameters0049{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			parameters0049{[]string{""}},
			[][]string{{""}},
		},
		{
			parameters0049{[]string{"a"}},
			[][]string{{"a"}},
		},
	}

	for _, test := range tests {
		t.Run("Test GroupAnagrams", func(t *testing.T) {
			// 完整輸入參數
			result := GroupAnagrams(test.parameters0049.strs)
			// compare 的方式需視情況調整
			utils.Sort2DArrayWithSortedSubArrays(test.ans)
			utils.Sort2DArrayWithSortedSubArrays(result)
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("GroupAnagrams(%+v) got %+v, want %+v", test.parameters0049, result, test.ans)
			}
		})
	}
}
