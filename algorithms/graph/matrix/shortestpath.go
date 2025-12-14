package matrix

// You are given an n x n binary matrix grid.
// Return the length of the shortest clear path in the matrix.
// If there is no clear path, return -1.

// A clear path in a binary matrix is a path from the top-left cell
// (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:

// All the visited cells of the path are 0.
// All the adjacent cells of the path are 8-directionally connected
// (i.e., they are different and they share an edge or a corner).
// The length of a clear path is the number of visited cells of this path.

// Example 1:

// Input: grid = [
//     [0,1,0],
//     [1,0,0],
//     [1,1,0]
// ]

// Output: 3
// Example 2:

// Input: grid = [
//     [1,0],
//     [1,1]
// ]

// Output: -1

// [
// 	[1,0,0],
// 	[1,1,0],
// 	[1,1,0]
// ]

type Pair struct {
	Row int
	Col int
}

var directions = []Pair{
	{-1, 0},  // up
	{0, 1},   // right
	{1, 0},   // down
	{0, -1},  // left
	{-1, 1},  // up-right
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, -1}, // up-left
}

// keep track of where we have visited
// continually explore nodes in a queue until finding bottom right OR queue becomes empty
func ShortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])

	start := Pair{0, 0}
	queue := []Pair{start}
	steps := 1


	// Could also be done as a map of Pair to bool or struct{}
	// Preallocating a slice like this is more efficient but less readable
	visited := make([][]bool, rows)
	for r := range rows {
		visited[r] = make([]bool, cols)
	}

	for len(queue) > 0 {
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				neighbour := Pair{
					Row: node.Row + direction.Row,
					Col: node.Col + direction.Col,
				}
				if isOutOfBounds(neighbour, rows, cols) ||
					grid[neighbour.Row][neighbour.Col] == 1 ||
					visited[neighbour.Row][neighbour.Col] {
					continue
				}

				if neighbour.Row == rows-1 && neighbour.Col == cols-1 {
					return steps + 1
				}

				visited[neighbour.Row][neighbour.Col] = true
				queue = append(queue, neighbour)
			}
		}
		steps++
	}

	return -1
}

func isOutOfBounds(pair Pair, rows, cols int) bool {
	if pair.Row < 0 || pair.Col < 0 {
		return true
	}

	if pair.Row > rows-1 || pair.Col > cols-1 {
		return true
	}

	return false
}
