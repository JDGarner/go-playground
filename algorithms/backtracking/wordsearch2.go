package backtracking

// Given an m x n board of characters and a list of strings words,
// return all words on the board.

// Each word must be constructed from letters of sequentially adjacent cells,
// where adjacent cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once in a word.

// Example 1:
// Input: board = [
// 	["o","a","a","n"],
// 	["e","t","a","e"],
// 	["i","h","k","r"],
// 	["i","f","l","v"]], words = ["oath","pea","eat","rain"]
// Output: ["eat","oath"]

type Direction struct {
	row int
	col int
}

var dirs = []Direction{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func findWords(board [][]byte, words []string) []string {
	numRows, numCols := len(board), len(board[0])

	foundWords := []string{}

	var dfs func(r, c, wordIndex int, word string, visited [][]bool) bool

	dfs = func(r, c, wordIndex int, word string, visited [][]bool) bool {
		// Base case: found the complete word
		if wordIndex == len(word)-1 && board[r][c] == word[wordIndex] {
			return true
		}

		if board[r][c] != word[wordIndex] {
			return false
		}

		visited[r][c] = true

		for _, dir := range dirs {
			neighbour := Direction{
				row: r + dir.row,
				col: c + dir.col,
			}

			if outOfBounds(neighbour, numRows, numCols) || visited[neighbour.row][neighbour.col] {
				continue
			}

			if dfs(neighbour.row, neighbour.col, wordIndex+1, word, visited) {
				return true
			}
		}

		visited[r][c] = false

		return false
	}

	// loop through the words and search for each one
	for _, word := range words {
	rowLoop:
		for r := range numRows {
			for c := range numCols {
				// Reset visited here?
				visited := make([][]bool, numRows)
				for r := range numRows {
					visited[r] = make([]bool, numCols)
				}

				if dfs(r, c, 0, word, visited) {
					foundWords = append(foundWords, word)
					break rowLoop
				}
			}
		}
	}

	return foundWords
}

func outOfBounds(pair Direction, rows, cols int) bool {
	if pair.row < 0 || pair.col < 0 {
		return true
	}

	if pair.row > rows-1 || pair.col > cols-1 {
		return true
	}

	return false
}
