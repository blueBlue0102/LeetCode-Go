package utils

import (
	"reflect"
	"testing"
)

func TestSort2DIntArray(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			// 測試空矩陣
			matrix: [][]int{},
			want:   [][]int{},
		},
		{
			// 測試只有一行的矩陣
			matrix: [][]int{{3, 2, 1}},
			want:   [][]int{{3, 2, 1}},
		},
		{
			// 測試多行多列的矩陣
			matrix: [][]int{{6, 5, 4}, {9, 8, 7}, {3, 2, 1}},
			want:   [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
		},
		{
			// 測試各行長度不同的矩陣
			matrix: [][]int{{3, 2, 1}, {9, 8, 7, 10}, {6, 5}},
			want:   [][]int{{3, 2, 1}, {6, 5}, {9, 8, 7, 10}},
		},
		{
			// 測試完全相反的排序順序
			matrix: [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}},
			want:   [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
		},
		{
			matrix: [][]int{{5}, {0}, {0, 1}, {5, 4, 3}, {5, 4, 2}, {5, 4, 2, 2}},
			want:   [][]int{{0}, {0, 1}, {5}, {5, 4, 2}, {5, 4, 2, 2}, {5, 4, 3}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Sort2DIntArray(test.matrix)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Sort2DIntArray(%v) = %v, want %v", test.matrix, got, test.want)
			}
		})
	}
}
