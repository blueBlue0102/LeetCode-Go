package utils

import (
	"reflect"
	"testing"
)

func TestSort2DIntArray(t *testing.T) {
	intTests := []struct {
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

	for _, test := range intTests {
		t.Run("", func(t *testing.T) {
			got := Sort2DArray(test.matrix)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Sort2DArray(%v) = %v, want %v", test.matrix, got, test.want)
			}
		})
	}
}

func TestSort2DStringArray(t *testing.T) {
	tests := []struct {
		matrix [][]string
		want   [][]string
	}{
		{
			// 測試空矩陣
			matrix: [][]string{},
			want:   [][]string{},
		},
		{
			// 測試只有一行的矩陣
			matrix: [][]string{{"c", "b", "a"}},
			want:   [][]string{{"c", "b", "a"}},
		},
		{
			// 測試多行多列的矩陣
			matrix: [][]string{{"f", "e", "d"}, {"i", "h", "g"}, {"c", "b", "a"}},
			want:   [][]string{{"c", "b", "a"}, {"f", "e", "d"}, {"i", "h", "g"}},
		},
		{
			// 測試各行長度不同的矩陣
			matrix: [][]string{{"c", "b", "a"}, {"i", "h", "g", "j"}, {"f", "e"}},
			want:   [][]string{{"c", "b", "a"}, {"f", "e"}, {"i", "h", "g", "j"}},
		},
		{
			// 測試完全相反的排序順序
			matrix: [][]string{{"i", "h", "g"}, {"f", "e", "d"}, {"c", "b", "a"}},
			want:   [][]string{{"c", "b", "a"}, {"f", "e", "d"}, {"i", "h", "g"}},
		},
		{
			matrix: [][]string{{"e"}, {"a"}, {"a", "b"}, {"e", "d", "c"}, {"e", "d", "b"}, {"e", "d", "b", "b"}},
			want:   [][]string{{"a"}, {"a", "b"}, {"e"}, {"e", "d", "b"}, {"e", "d", "b", "b"}, {"e", "d", "c"}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Sort2DArray(test.matrix)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Sort2DArray(%v) = %v, want %v", test.matrix, got, test.want)
			}
		})
	}
}

func TestSort2DArrayWithSortedSubArrays(t *testing.T) {
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
			matrix: [][]int{{3, 1, 2}},
			want:   [][]int{{1, 2, 3}},
		},
		{
			// 測試多行多列的矩陣
			matrix: [][]int{{6, 4, 5}, {9, 7, 8}, {3, 1, 2}},
			want:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			// 測試各行長度不同的矩陣
			matrix: [][]int{{3, 1, 2}, {8, 9, 7, 10}, {5, 4}},
			want:   [][]int{{1, 2, 3}, {4, 5}, {7, 8, 9, 10}},
		},
		{
			// 測試完全相反的排序順序
			matrix: [][]int{{9, 7, 8}, {6, 4, 5}, {3, 1, 2}},
			want:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			// 測試包含重複元素的矩陣
			matrix: [][]int{{5}, {1}, {1, 2}, {5, 4, 3}, {5, 4, 2}, {5, 4, 2, 2}},
			want:   [][]int{{1}, {1, 2}, {2, 2, 4, 5}, {2, 4, 5}, {3, 4, 5}, {5}},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Sort2DArrayWithSortedSubArrays(test.matrix)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Sort2DArrayWithSortedSubArrays(%v) = %v, want %v", test.matrix, got, test.want)
			}
		})
	}
}
