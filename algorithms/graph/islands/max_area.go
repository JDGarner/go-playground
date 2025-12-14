package islands

// You are given a matrix grid where grid[i] is either a 0 (representing water)
// or 1 (representing land).

// An island is defined as a group of 1's connected horizontally or vertically.
// You may assume all four edges of the grid are surrounded by water.

// The area of an island is defined as the number of cells within the island.
// Return the maximum area of an island in grid. If no island exists, return 0.

// Input: grid = [
//   [0,1,1,0,1],
//   [1,0,1,0,1],
//   [0,1,1,0,1],
//   [0,1,0,0,1]
// ]
// Output: 6

var directions2 = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func MaxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	largest := 0
	var islandSize func(r, c int) int

	islandSize = func(r, c int) int {
		size := 1
		grid[r][c] = 0

		for _, pair := range directions2 {
			rPair := r + pair[0]
			cPair := c + pair[1]

			if rPair < 0 || cPair < 0 || rPair > rows-1 || cPair > cols-1 || grid[rPair][cPair] == 0 {
				continue
			}
			size += islandSize(rPair, cPair)
		}

		return size
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				size := islandSize(i, j)
				if size > largest {
					largest = size
				}
			}
		}
	}

	return largest
}
