package intervals

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		toBeRemoved []int
		expected    [][]int
	}{
		{
			name:        "Example 1: Remove middle portion across multiple intervals",
			intervals:   [][]int{{0, 2}, {3, 4}, {5, 7}},
			toBeRemoved: []int{1, 6},
			expected:    [][]int{{0, 1}, {6, 7}},
		},
		{
			name:        "Example 2: Remove middle of single interval",
			intervals:   [][]int{{0, 5}},
			toBeRemoved: []int{2, 3},
			expected:    [][]int{{0, 2}, {3, 5}},
		},
		{
			name:        "Example 3: Remove across negative and positive intervals",
			intervals:   [][]int{{-5, -4}, {-3, -2}, {1, 2}, {3, 5}, {8, 9}},
			toBeRemoved: []int{-1, 4},
			expected:    [][]int{{-5, -4}, {-3, -2}, {4, 5}, {8, 9}},
		},
		{
			name:        "Remove completely before all intervals",
			intervals:   [][]int{{5, 7}, {10, 12}},
			toBeRemoved: []int{1, 3},
			expected:    [][]int{{5, 7}, {10, 12}},
		},
		{
			name:        "Remove completely after all intervals",
			intervals:   [][]int{{1, 3}, {5, 7}},
			toBeRemoved: []int{10, 12},
			expected:    [][]int{{1, 3}, {5, 7}},
		},
		{
			name:        "Remove entire single interval exactly",
			intervals:   [][]int{{5, 10}},
			toBeRemoved: []int{5, 10},
			expected:    [][]int{},
		},
		{
			name:        "Remove all intervals completely",
			intervals:   [][]int{{1, 3}, {5, 7}, {9, 11}},
			toBeRemoved: []int{0, 15},
			expected:    [][]int{},
		},
		{
			name:        "Remove overlaps left edge of first interval",
			intervals:   [][]int{{5, 10}, {15, 20}},
			toBeRemoved: []int{3, 7},
			expected:    [][]int{{7, 10}, {15, 20}},
		},
		{
			name:        "Remove overlaps right edge of last interval",
			intervals:   [][]int{{5, 10}, {15, 20}},
			toBeRemoved: []int{18, 25},
			expected:    [][]int{{5, 10}, {15, 18}},
		},
		{
			name:        "Remove touches left boundary exactly",
			intervals:   [][]int{{5, 10}},
			toBeRemoved: []int{5, 7},
			expected:    [][]int{{7, 10}},
		},
		{
			name:        "Remove touches right boundary exactly",
			intervals:   [][]int{{5, 10}},
			toBeRemoved: []int{7, 10},
			expected:    [][]int{{5, 7}},
		},
		{
			name:        "Remove between intervals with no overlap",
			intervals:   [][]int{{1, 3}, {7, 9}},
			toBeRemoved: []int{4, 6},
			expected:    [][]int{{1, 3}, {7, 9}},
		},
		{
			name:        "Single interval, remove from left edge to middle",
			intervals:   [][]int{{0, 10}},
			toBeRemoved: []int{0, 5},
			expected:    [][]int{{5, 10}},
		},
		{
			name:        "Single interval, remove from middle to right edge",
			intervals:   [][]int{{0, 10}},
			toBeRemoved: []int{5, 10},
			expected:    [][]int{{0, 5}},
		},
		{
			name:        "Remove spans gap between two intervals",
			intervals:   [][]int{{1, 5}, {10, 15}},
			toBeRemoved: []int{3, 12},
			expected:    [][]int{{1, 3}, {12, 15}},
		},
		{
			name:        "Multiple intervals, remove middle one completely",
			intervals:   [][]int{{1, 3}, {5, 7}, {9, 11}},
			toBeRemoved: []int{5, 7},
			expected:    [][]int{{1, 3}, {9, 11}},
		},
		{
			name:        "Empty intervals list",
			intervals:   [][]int{},
			toBeRemoved: []int{5, 10},
			expected:    [][]int{},
		},
		{
			name:        "Single point interval, remove overlaps",
			intervals:   [][]int{{5, 5}},
			toBeRemoved: []int{3, 7},
			expected:    [][]int{},
		},
		{
			name:        "Multiple intervals with gaps, remove overlaps first and third",
			intervals:   [][]int{{1, 4}, {6, 8}, {10, 15}},
			toBeRemoved: []int{2, 12},
			expected:    [][]int{{1, 2}, {12, 15}},
		},
		{
			name:        "Remove exactly matches interval boundaries",
			intervals:   [][]int{{1, 5}, {5, 10}, {10, 15}},
			toBeRemoved: []int{5, 10},
			expected:    [][]int{{1, 5}, {10, 15}},
		},
		{
			name:        "Negative intervals only",
			intervals:   [][]int{{-10, -5}, {-3, -1}},
			toBeRemoved: []int{-7, -2},
			expected:    [][]int{{-10, -7}, {-2, -1}},
		},
		{
			name:        "Remove spans from negative to positive",
			intervals:   [][]int{{-5, -1}, {0, 5}, {6, 10}},
			toBeRemoved: []int{-3, 7},
			expected:    [][]int{{-5, -3}, {7, 10}},
		},
		{
			name:        "Large intervals with small removal",
			intervals:   [][]int{{0, 1000}},
			toBeRemoved: []int{400, 600},
			expected:    [][]int{{0, 400}, {600, 1000}},
		},
		{
			name:        "Remove at exact start of first interval",
			intervals:   [][]int{{10, 20}, {30, 40}},
			toBeRemoved: []int{10, 15},
			expected:    [][]int{{15, 20}, {30, 40}},
		},
		{
			name:        "Remove at exact end of last interval",
			intervals:   [][]int{{10, 20}, {30, 40}},
			toBeRemoved: []int{35, 40},
			expected:    [][]int{{10, 20}, {30, 35}},
		},
		{
			name:        "Consecutive intervals, remove affects multiple",
			intervals:   [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			toBeRemoved: []int{2, 4},
			expected:    [][]int{{1, 2}, {4, 5}},
		},
		{
			name:        "Remove completely contains multiple intervals",
			intervals:   [][]int{{2, 4}, {6, 8}, {10, 12}},
			toBeRemoved: []int{1, 13},
			expected:    [][]int{},
		},
		{
			name:        "Remove partially overlaps first and completely contains rest",
			intervals:   [][]int{{1, 5}, {7, 10}, {12, 15}},
			toBeRemoved: []int{3, 20},
			expected:    [][]int{{1, 3}},
		},
		{
			name:        "Single interval where remove is completely inside",
			intervals:   [][]int{{0, 100}},
			toBeRemoved: []int{25, 75},
			expected:    [][]int{{0, 25}, {75, 100}},
		},
		{
			name:        "Zero-width intervals",
			intervals:   [][]int{{5, 5}, {10, 10}, {15, 15}},
			toBeRemoved: []int{10, 10},
			expected:    [][]int{{5, 5}, {15, 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := remove(tt.intervals, tt.toBeRemoved)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("remove(%v, %v) = %v; expected %v",
					tt.intervals, tt.toBeRemoved, result, tt.expected)
			}
		})
	}
}

func TestRemoveEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		toBeRemoved []int
		expected    [][]int
	}{
		{
			name:        "Remove with same start and end",
			intervals:   [][]int{{1, 10}},
			toBeRemoved: []int{5, 5},
			expected:    [][]int{{1, 5}, {5, 10}},
		},
		{
			name:        "Intervals touching but not overlapping",
			intervals:   [][]int{{1, 5}, {5, 10}},
			toBeRemoved: []int{3, 7},
			expected:    [][]int{{1, 3}, {7, 10}},
		},
		{
			name:        "Many small intervals with larger removal",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			toBeRemoved: []int{2, 9},
			expected:    [][]int{{1, 2}, {9, 10}},
		},
		{
			name:        "Very large numbers",
			intervals:   [][]int{{1000000, 2000000}, {3000000, 4000000}},
			toBeRemoved: []int{1500000, 3500000},
			expected:    [][]int{{1000000, 1500000}, {3500000, 4000000}},
		},
		{
			name:        "Very small negative numbers",
			intervals:   [][]int{{-1000000, -500000}, {-400000, -100000}},
			toBeRemoved: []int{-600000, -300000},
			expected:    [][]int{{-1000000, -600000}, {-300000, -100000}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := remove(tt.intervals, tt.toBeRemoved)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("remove(%v, %v) = %v; expected %v",
					tt.intervals, tt.toBeRemoved, result, tt.expected)
			}
		})
	}
}
