package rottenfruit_test

import (
	"testing"

	"github.com/JDGarner/go-playground/algorithms/graph/rottenfruit"
	"github.com/JDGarner/go-playground/datastructures/graph/matrix"
)

func TestRottenFruit(t *testing.T) {
	t.Run("rotten fruit problem - case 1", func(t *testing.T) {
		input := matrix.New([][]int{
			{1, 1, 0},
			{0, 1, 1},
			{0, 1, 2},
		})
		expected := 4

		actual := rottenfruit.RunRottenFruit(input)

		if actual != expected {
			t.Fatalf("output not expected: got %d, want %d", actual, expected)
		}
	})

	t.Run("rotten fruit problem - case 2", func(t *testing.T) {
		input := matrix.New([][]int{
			{1, 0, 1},
			{0, 2, 0},
			{1, 0, 1},
		})
		expected := -1

		actual := rottenfruit.RunRottenFruit(input)

		if actual != expected {
			t.Fatalf("output not expected: got %d, want %d", actual, expected)
		}
	})
}
