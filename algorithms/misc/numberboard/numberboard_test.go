package numberboard

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(xs []int) []int {
	sort.Ints(xs)
	return xs
}

func TestReachablePositions(t *testing.T) {
	tests := []struct {
		name        string
		start       int
		diceSides   int
		teleporters map[int]int
		expected    []int
	}{
		{
			name:        "no teleporters",
			start:       1,
			diceSides:   3,
			teleporters: nil,
			expected:    []int{2, 3, 4},
		},
		{
			name:      "single teleporter",
			start:     5,
			diceSides: 3,
			teleporters: map[int]int{
				7: 12,
			},
			// roll=1 -> 6
			// roll=2 -> 7 -> 12
			// roll=3 -> 8
			expected: []int{6, 7, 12, 8},
		},
		{
			name:      "teleporter chain",
			start:     1,
			diceSides: 1,
			teleporters: map[int]int{
				2: 5,
				5: 9,
			},
			// roll=1 -> 2 -> 5 -> 9
			expected: []int{2, 5, 9},
		},
		{
			name:      "multiple teleporters across rolls",
			start:     5,
			diceSides: 4,
			teleporters: map[int]int{
				7: 12,
				9: 3,
			},
			// roll=1 -> 6
			// roll=2 -> 7 -> 12
			// roll=3 -> 8
			// roll=4 -> 9 -> 3
			expected: []int{6, 7, 12, 8, 9, 3},
		},
		{
			name:      "teleporter cycle",
			start:     1,
			diceSides: 2,
			teleporters: map[int]int{
				2: 6,
				6: 3,
				3: 2,
			},
			// roll=1 -> 2 -> 6 -> 3 -> cycle
			// roll=2 -> 3 -> 2 -> 6 -> cycle
			expected: []int{2, 3, 6},
		},
		{
			name:      "overlapping results from different rolls",
			start:     0,
			diceSides: 3,
			teleporters: map[int]int{
				1: 5,
				2: 5,
			},
			// roll=1 -> 1 -> 5
			// roll=2 -> 2 -> 5
			// roll=3 -> 3
			expected: []int{1, 2, 3, 5},
		},
		{
			name:        "zero dice sides",
			start:       10,
			diceSides:   0,
			teleporters: nil,
			expected:    []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReachablePositions(tt.start, tt.diceSides, tt.teleporters)

			if !reflect.DeepEqual(
				normalize(result),
				normalize(tt.expected),
			) {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
