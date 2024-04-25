package leetcode

import (
	"reflect"
	"testing"
)

// 填入 function input type
type parameters0690 struct {
	employees []*Employee
	id        int
}

func TestEmployeeImportance(t *testing.T) {
	tests := []struct {
		parameters0690
		// 填入 function output type
		ans int
	}{
		// 填入 test case
		{
			parameters0690{[]*Employee{{1, 5, []int{2, 3}}, {2, 3, []int{}}, {3, 3, []int{}}}, 1},
			11,
		},
		{
			parameters0690{[]*Employee{{1, 2, []int{5}}, {5, -3, []int{}}}, 5},
			-3,
		},
	}

	for _, test := range tests {
		t.Run("Test EmployeeImportance", func(t *testing.T) {
			// 完整輸入參數
			result := EmployeeImportance(test.parameters0690.employees, test.parameters0690.id)
			// compare 的方式需視情況調整
			if !reflect.DeepEqual(result, test.ans) {
				t.Errorf("EmployeeImportance(%+v) got %+v, want %+v", test.parameters0690, result, test.ans)
			}
		})
	}
}
