package intervals

import (
	"reflect"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		lower int
		upper int
		want  [][]int
	}{
		// Basic examples from problem statement
		{
			name:  "example 1 - multiple gaps",
			nums:  []int{0, 1, 3, 50, 75},
			lower: 0,
			upper: 99,
			want:  [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}},
		},
		{
			name:  "example 2 - no missing ranges",
			nums:  []int{-1},
			lower: -1,
			upper: -1,
			want:  [][]int{},
		},

		// Empty array cases
		{
			name:  "empty array - single element range",
			nums:  []int{},
			lower: 1,
			upper: 1,
			want:  [][]int{{1, 1}},
		},
		{
			name:  "empty array - multiple elements",
			nums:  []int{},
			lower: 1,
			upper: 5,
			want:  [][]int{{1, 5}},
		},
		{
			name:  "empty array - large range",
			nums:  []int{},
			lower: -100,
			upper: 100,
			want:  [][]int{{-100, 100}},
		},

		// Single element array cases
		{
			name:  "single element - at lower bound",
			nums:  []int{0},
			lower: 0,
			upper: 5,
			want:  [][]int{{1, 5}},
		},
		{
			name:  "single element - at upper bound",
			nums:  []int{5},
			lower: 0,
			upper: 5,
			want:  [][]int{{0, 4}},
		},
		{
			name:  "single element - in middle",
			nums:  []int{3},
			lower: 0,
			upper: 5,
			want:  [][]int{{0, 2}, {4, 5}},
		},
		{
			name:  "single element - fills entire range",
			nums:  []int{5},
			lower: 5,
			upper: 5,
			want:  [][]int{},
		},

		// Consecutive elements
		{
			name:  "all consecutive - no gaps",
			nums:  []int{1, 2, 3, 4, 5},
			lower: 1,
			upper: 5,
			want:  [][]int{},
		},
		{
			name:  "consecutive with gap at start",
			nums:  []int{3, 4, 5},
			lower: 1,
			upper: 5,
			want:  [][]int{{1, 2}},
		},
		{
			name:  "consecutive with gap at end",
			nums:  []int{1, 2, 3},
			lower: 1,
			upper: 5,
			want:  [][]int{{4, 5}},
		},

		// Single element gaps
		{
			name:  "multiple single element gaps",
			nums:  []int{1, 3, 5, 7, 9},
			lower: 1,
			upper: 9,
			want:  [][]int{{2, 2}, {4, 4}, {6, 6}, {8, 8}},
		},
		{
			name:  "single element gap at boundaries",
			nums:  []int{2, 4},
			lower: 1,
			upper: 5,
			want:  [][]int{{1, 1}, {3, 3}, {5, 5}},
		},

		// Negative numbers
		{
			name:  "negative range - all missing",
			nums:  []int{},
			lower: -5,
			upper: -1,
			want:  [][]int{{-5, -1}},
		},
		{
			name:  "negative range - with gaps",
			nums:  []int{-5, -3, -1},
			lower: -5,
			upper: -1,
			want:  [][]int{{-4, -4}, {-2, -2}},
		},
		{
			name:  "negative to positive range",
			nums:  []int{-2, 0, 2},
			lower: -3,
			upper: 3,
			want:  [][]int{{-3, -3}, {-1, -1}, {1, 1}, {3, 3}},
		},

		// Large gaps
		{
			name:  "very large gap",
			nums:  []int{0, 1000000},
			lower: 0,
			upper: 1000000,
			want:  [][]int{{1, 999999}},
		},
		{
			name:  "multiple large gaps",
			nums:  []int{0, 100, 200},
			lower: 0,
			upper: 300,
			want:  [][]int{{1, 99}, {101, 199}, {201, 300}},
		},

		// Edge cases with boundaries
		{
			name:  "nums covers exact boundaries",
			nums:  []int{1, 10},
			lower: 1,
			upper: 10,
			want:  [][]int{{2, 9}},
		},
		{
			name:  "nums at both boundaries - no middle",
			nums:  []int{1, 2},
			lower: 1,
			upper: 2,
			want:  [][]int{},
		},

		// Two element arrays
		{
			name:  "two elements - consecutive",
			nums:  []int{5, 6},
			lower: 1,
			upper: 10,
			want:  [][]int{{1, 4}, {7, 10}},
		},
		{
			name:  "two elements - with gap",
			nums:  []int{3, 7},
			lower: 1,
			upper: 10,
			want:  [][]int{{1, 2}, {4, 6}, {8, 10}},
		},

		// Maximum density
		{
			name:  "almost full - one missing at start",
			nums:  []int{2, 3, 4, 5},
			lower: 1,
			upper: 5,
			want:  [][]int{{1, 1}},
		},
		{
			name:  "almost full - one missing at end",
			nums:  []int{1, 2, 3, 4},
			lower: 1,
			upper: 5,
			want:  [][]int{{5, 5}},
		},
		{
			name:  "almost full - one missing in middle",
			nums:  []int{1, 2, 4, 5},
			lower: 1,
			upper: 5,
			want:  [][]int{{3, 3}},
		},

		// Zero in range
		{
			name:  "zero at lower bound",
			nums:  []int{0, 2, 4},
			lower: 0,
			upper: 5,
			want:  [][]int{{1, 1}, {3, 3}, {5, 5}},
		},
		{
			name:  "zero in middle",
			nums:  []int{-2, 0, 2},
			lower: -2,
			upper: 2,
			want:  [][]int{{-1, -1}, {1, 1}},
		},

		// Sparse arrays
		{
			name:  "very sparse array",
			nums:  []int{10, 30, 50, 70, 90},
			lower: 0,
			upper: 100,
			want:  [][]int{{0, 9}, {11, 29}, {31, 49}, {51, 69}, {71, 89}, {91, 100}},
		},

		// Same lower and upper
		{
			name:  "range of size 1 - empty",
			nums:  []int{},
			lower: 42,
			upper: 42,
			want:  [][]int{{42, 42}},
		},
		{
			name:  "range of size 1 - filled",
			nums:  []int{42},
			lower: 42,
			upper: 42,
			want:  [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMissingRanges(tt.nums, tt.lower, tt.upper)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test that the function returns a non-nil slice even when empty
func TestFindMissingRanges_ReturnsNonNilSlice(t *testing.T) {
	result := findMissingRanges([]int{1, 2, 3}, 1, 3)
	if result == nil {
		t.Error("Expected non-nil slice, got nil")
	}
}

// Benchmark for performance testing
func BenchmarkFindMissingRanges(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i * 2 // Every even number
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMissingRanges(nums, 0, 2000)
	}
}
