package backtracking

// Given a 2-D grid of characters board and a string word,
// return true if the word is present in the grid, otherwise return false.

// For the word to be present it must be possible to form it with a path
// in the board with horizontally or vertically neighboring cells.
// The same cell may not be used more than once in a word.

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

func exist(board [][]byte, word string) bool {
	numRows, numCols := len(board), len(board[0])
	visited := make([][]bool, numRows)
	for i := range visited {
		visited[i] = make([]bool, numCols)
	}

	var dfs func(node Pair, index int) bool

	dfs = func(node Pair, index int) bool {
		// Base case: found the complete word
		if index == len(word) {
			return true
		}

		if isOutOfBounds(node, numRows, numCols) || visited[node.row][node.col] {
			return false
		}

		// Check if current cell matches the current character
		if board[node.row][node.col] != word[index] {
			return false
		}

		visited[node.row][node.col] = true

		for _, direction := range directions {
			neighbour := Pair{
				row: node.row + direction.row,
				col: node.col + direction.col,
			}
			if dfs(neighbour, index+1) {
				return true
			}
		}

		// Backtrack: unmark as visited
		visited[node.row][node.col] = false

		return false
	}

	for r := range numRows {
		for c := range numCols {
			node := Pair{r, c}

			if dfs(node, 0) {
				return true
			}
		}
	}

	return false
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
