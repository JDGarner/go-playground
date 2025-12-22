package regions

// You are given an m x n matrix board containing letters 'X' and 'O',
// capture regions that are surrounded:

// Connect: A cell is connected to adjacent cells horizontally or vertically.

// Region: To form a region connect every 'O' cell.

// Surround: The region is surrounded with 'X' cells if you can connect the
// region with 'X' cells and none of the region cells are on the edge of the board.

// To capture a surrounded region, replace all 'O's with 'X's in-place within
// the original board. You do not need to return anything.

// Example:
// Input:
// [
// 	["X","X","X","X"],
// 	["X","O","O","X"],
// 	["X","X","O","X"],
// 	["X","O","X","X"]
// ]

// Output:
// [
// 	["X","X","X","X"],
// 	["X","X","X","X"],
// 	["X","X","X","X"],
// 	["X","O","X","X"]
// ]

type Pair struct {
	row int
	col int
}

var directions = []Pair{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func solve(board [][]byte) {
	// for each "0" is there a path that goes to the edge of the board
	// - if no - capture any visited 0s
	// - if yes - just mark them as visited

	// Solution 2:
	// Any 0 reachable from the edges cannot be captured - mark all those as 'R'
	// Turn all remaining '0' into 'X'
	// Turn all the 'R' back to '0'

	numRows, numCols := len(board), len(board[0])

	var markReachable func(r, c int)

	markReachable = func(r, c int) {
		if board[r][c] == 'X' || board[r][c] == 'R' {
			return
		}

		// Must be a 0 that cannot be captured
		board[r][c] = 'R'

		for _, direction := range directions {
			neighbour := Pair{r + direction.row, c + direction.col}

			if isOutOfBounds(neighbour, numRows, numCols) {
				continue
			}

			markReachable(neighbour.row, neighbour.col)
		}
	}

	// left, right column
	for r := range numRows {
		markReachable(r, 0)
		markReachable(r, numCols-1)
	}

	// top, bottom row
	for c := range numCols {
		markReachable(0, c)
		markReachable(numRows-1, c)
	}

	// Turn all captured to X
	for r := range numRows {
		for c := range numCols {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}

	// Turn all reachable back to 0
	for r := range numRows {
		for c := range numCols {
			if board[r][c] == 'R' {
				board[r][c] = 'O'
			}
		}
	}
}

func isOutOfBounds(pair Pair, rows, cols int) bool {
	if pair.row < 0 || pair.col < 0 {
		return true
	}

	if pair.row > rows-1 || pair.col > cols-1 {
		return true
	}

	return false
}
