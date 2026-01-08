package snakegame

import (
	"testing"
)

// TestBasicExample tests the example from the problem description
func TestBasicExample(t *testing.T) {
	game := NewSnakeGame(3, 2, [][]int{{1, 2}, {0, 1}})

	tests := []struct {
		direction string
		expected  int
	}{
		{"R", 0},
		{"D", 0},
		{"R", 1},
		{"U", 1},
		{"L", 2},
		{"U", -1},
	}

	for i, tt := range tests {
		result := game.move(tt.direction)
		if result != tt.expected {
			t.Errorf("Move %d: expected %d, got %d", i+1, tt.expected, result)
		}
	}
}

// TestInitialState tests the game's initial state
func TestInitialState(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{2, 2}})

	// Snake should be at (0,0) initially
	// Moving right should work
	result := game.move("R")
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// TestOutOfBoundsTop tests hitting the top wall
func TestOutOfBoundsTop(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{2, 2}})

	// Start at (0,0), try to move up
	result := game.move("U")
	if result != -1 {
		t.Errorf("Expected -1 for out of bounds, got %d", result)
	}
}

// TestOutOfBoundsLeft tests hitting the left wall
func TestOutOfBoundsLeft(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{2, 2}})

	// Start at (0,0), try to move left
	result := game.move("L")
	if result != -1 {
		t.Errorf("Expected -1 for out of bounds, got %d", result)
	}
}

// TestOutOfBoundsRight tests hitting the right wall
func TestOutOfBoundsRight(t *testing.T) {
	game := NewSnakeGame(3, 3, [][]int{{1, 1}})

	// Move right to edge
	game.move("R")           // (0,1)
	game.move("R")           // (0,2)
	result := game.move("R") // Out of bounds

	if result != -1 {
		t.Errorf("Expected -1 for out of bounds, got %d", result)
	}
}

// TestOutOfBoundsBottom tests hitting the bottom wall
func TestOutOfBoundsBottom(t *testing.T) {
	game := NewSnakeGame(3, 3, [][]int{{1, 1}})

	// Move down to edge
	game.move("D")           // (1,0)
	game.move("D")           // (2,0)
	result := game.move("D") // Out of bounds

	if result != -1 {
		t.Errorf("Expected -1 for out of bounds, got %d", result)
	}
}

// TestEatingFood tests that eating food increases score
func TestEatingFood(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{0, 1}, {0, 2}})

	// Move right to eat first food at (0,1)
	result := game.move("R")
	if result != 1 {
		t.Errorf("Expected score 1 after eating food, got %d", result)
	}

	// Move right again to eat second food at (0,2)
	result = game.move("R")
	if result != 2 {
		t.Errorf("Expected score 2 after eating food, got %d", result)
	}
}

// TestSnakeGrowth tests that snake grows after eating food
func TestSnakeGrowth(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{0, 2}})

	// Move right twice to reach food at (0,2)
	game.move("R") // (0,1)
	game.move("R") // (0,2) - eat food, length is now 2

	// Now snake occupies (0,2) and (0,1)
	// Move down
	game.move("D") // head at (1,2), tail at (0,2)

	// snakegame_test.go:118: snake body 2:  [{0 1} {0 2}]
	// snakegame_test.go:123: snake body 3:  [{0 2} {1 0}]

	// Move left - should not hit body
	result := game.move("L") // head at (1,1), body at (1,2) and (0,2)
	if result == -1 {
		t.Errorf("Snake should not collide with itself here")
	}
}

// TestSelfCollision tests snake running into itself
func TestSelfCollision(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{0, 1}, {0, 2}, {1, 2}})

	// Eat first food
	game.move("R") // (0,1) - eat food, length 2
	// Eat second food
	game.move("R") // (0,2) - eat food, length 3
	// Eat third food
	game.move("D") // (1,2) - eat food, length 4

	// Now snake is: head (1,2), body (0,2), (0,1), (0,0)
	game.move("L") // (1,1)
	// Snake is: head (1,1), body (1,2), (0,2), (0,1)

	result := game.move("U") // (0,1) - collision with body
	if result != -1 {
		t.Errorf("Expected -1 for self collision, got %d", result)
	}
}

// TestNoFood tests game with no food
func TestNoFood(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{})

	result := game.move("R")
	if result != 0 {
		t.Errorf("Expected score 0 with no food, got %d", result)
	}

	result = game.move("D")
	if result != 0 {
		t.Errorf("Expected score 0 with no food, got %d", result)
	}
}

// TestMinimalGrid tests 1x1 grid
func TestMinimalGrid(t *testing.T) {
	game := NewSnakeGame(1, 1, [][]int{})

	// Any move should be out of bounds
	result := game.move("R")
	if result != -1 {
		t.Errorf("Expected -1 for 1x1 grid, got %d", result)
	}
}

// TestLongSnake tests a snake that grows very long
func TestLongSnake(t *testing.T) {
	food := [][]int{}
	for i := 1; i < 10; i++ {
		food = append(food, []int{0, i})
	}

	game := NewSnakeGame(15, 5, food)

	// Eat 9 pieces of food
	for i := 0; i < 9; i++ {
		result := game.move("R")
		if result != i+1 {
			t.Errorf("Move %d: expected score %d, got %d", i+1, i+1, result)
		}
	}
}

// TestAlternatingDirections tests moving in different directions
func TestAlternatingDirections(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{2, 0}})

	moves := []struct {
		direction string
		expected  int
	}{
		{"D", 0}, // (1,0)
		{"D", 1}, // (2,0) - eat food
		{"R", 1}, // (2,1)
		{"U", 1}, // (1,1)
		{"L", 1}, // (1,0)
	}

	for i, move := range moves {
		result := game.move(move.direction)
		if result != move.expected {
			t.Errorf("Move %d: expected %d, got %d", i+1, move.expected, result)
		}
	}
}

// TestMultipleGamesIndependence tests that multiple game instances are independent
func TestMultipleGamesIndependence(t *testing.T) {
	game1 := NewSnakeGame(3, 3, [][]int{{0, 1}})
	game2 := NewSnakeGame(5, 5, [][]int{{0, 2}})

	result1 := game1.move("R")
	result2 := game2.move("R")

	if result1 != 1 {
		t.Errorf("Game1 expected score 1, got %d", result1)
	}

	if result2 != 0 {
		t.Errorf("Game2 expected score 0, got %d", result2)
	}
}

// TestGameOverPersists tests that after game over, all moves return -1
func TestGameOverPersists(t *testing.T) {
	game := NewSnakeGame(3, 3, [][]int{{1, 1}})

	// Cause game over
	game.move("U") // Out of bounds

	// Try more moves
	result := game.move("R")
	if result != -1 {
		t.Errorf("Expected -1 after game over, got %d", result)
	}

	result = game.move("D")
	if result != -1 {
		t.Errorf("Expected -1 after game over, got %d", result)
	}
}

// TestComplexPath tests a complex movement pattern
func TestComplexPath(t *testing.T) {
	game := NewSnakeGame(10, 10, [][]int{{2, 1}, {3, 1}, {3, 2}})

	moves := []struct {
		direction string
		expected  int
	}{
		{"D", 0}, // (1,0)
		{"D", 0}, // (2,0)
		{"R", 1}, // (2,1) - eat food, length 2
		{"D", 2}, // (3,1) - eat food, length 3
		{"R", 3}, // (3,2) - eat food, length 4
	}

	for i, move := range moves {
		result := game.move(move.direction)
		if result != move.expected {
			t.Errorf("Move %d (%s): expected %d, got %d", i+1, move.direction, move.expected, result)
		}
	}
}

// TestSnakeCannotReverseIntoItself tests immediate reversal after growing
func TestSnakeCannotReverseIntoItself(t *testing.T) {
	game := NewSnakeGame(5, 5, [][]int{{0, 1}})

	game.move("R") // (0,1) - eat food, length 2, body at (0,0)

	// Try to move left immediately (into body)
	result := game.move("L") // Should hit body at (0,0)
	if result != -1 {
		t.Errorf("Expected -1 for reversing into body, got %d", result)
	}
}

// TestSelfCollisionBugWithTailMovement tests the critical bug where
// collision detection incorrectly checks the tail position that will move away
func TestSelfCollisionBugWithTailMovement(t *testing.T) {
	// This test specifically targets the bug in collidedWithSelf
	// The snake should be able to move to where its tail WAS, since the tail moves
	game := NewSnakeGame(5, 5, [][]int{})

	// Create a scenario where snake will move in a square
	// Start at (0,0)
	game.move("R") // Snake at (0,1), tail gone from (0,0)
	game.move("D") // Snake at (1,1), tail gone from (0,1)
	game.move("L") // Snake at (1,0), tail gone from (1,1)

	// Now try to move up to (0,0) - this should be VALID
	// The snake's body is currently only at (1,0)
	// The tail position (1,0) will move away when we move up
	result := game.move("U") // Should move to (0,0) successfully

	if result == -1 {
		t.Errorf("BUG DETECTED: Snake should be able to move to where tail was. Got -1, expected 0")
		t.Errorf("This indicates collidedWithSelf is checking the tail position incorrectly")
	}
	if result != 0 {
		t.Errorf("Expected score 0, got %d", result)
	}
}

// TestSelfCollisionWithGrowth tests collision detection when snake is growing
func TestSelfCollisionWithGrowth(t *testing.T) {
	// When snake eats food and grows, the tail does NOT move
	// So collision detection must check the entire body including tail
	game := NewSnakeGame(5, 5, [][]int{{0, 1}, {1, 1}, {1, 0}})

	// Eat first food
	game.move("R") // (0,1) - eat food, length 2: [(0,0), (0,1)]

	// Eat second food
	game.move("D") // (1,1) - eat food, length 3: [(0,0), (0,1), (1,1)]

	// Eat third food
	game.move("L") // (1,0) - eat food, length 4: [(0,0), (0,1), (1,1), (1,0)]

	// Now try to move up to (0,0) - this should FAIL
	// The tail at (0,0) did NOT move because we just ate food
	result := game.move("U") // Should collide with body at (0,0)

	if result != -1 {
		t.Errorf("Expected -1 for collision with tail after growth, got %d", result)
		t.Errorf("When snake grows, tail doesn't move, so (0,0) should cause collision")
	}
}

// TestFalsePositiveCollision tests a specific false positive scenario
func TestFalsePositiveCollision(t *testing.T) {
	// This creates a longer path to demonstrate the bug more clearly
	game := NewSnakeGame(10, 10, [][]int{})

	// Move in a pattern: Right, Down, Down, Left, Up, Up
	// This creates an 'S' shape where we revisit old positions
	moves := []struct {
		direction string
		expected  int
		position  string
	}{
		{"R", 0, "(0,1)"},
		{"R", 0, "(0,2)"},
		{"D", 0, "(1,2)"},
		{"D", 0, "(2,2)"},
		{"L", 0, "(2,1)"},
		{"L", 0, "(2,0)"},
		{"U", 0, "(1,0)"}, // Snake body is only at (2,0), tail will move
		{"U", 0, "(0,0)"}, // Snake body is only at (1,0), tail will move - should succeed
	}

	for i, move := range moves {
		result := game.move(move.direction)
		if result != move.expected {
			t.Errorf("Move %d to %s: expected %d, got %d",
				i+1, move.position, move.expected, result)
			if result == -1 {
				t.Errorf("FALSE POSITIVE: Collision detected when snake should be able to move")
			}
		}
	}
}
