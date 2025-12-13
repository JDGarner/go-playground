package islands

// Given a 2D grid grid where '1' represents land and '0' represents water,
// count and return the number of islands.

// An island is formed by connecting adjacent lands horizontally or vertically
// and is surrounded by water.
// You may assume water is surrounding the grid (i.e., all the edges are water).

// Example 1:

// Input: grid = [
//     ["0","1","1","1","0"],
//     ["0","1","0","1","0"],
//     ["1","1","0","0","0"],
//     ["0","0","0","0","0"]
//   ]
// Output: 1
// Example 2:

// Input: grid = [
//     ["1","1","0","0","1"],
//     ["1","1","0","0","1"],
//     ["0","0","1","0","0"],
//     ["0","0","0","1","1"]
//   ]
// Output: 4

type Pair struct {
	Row int
	Col int
}

var directions = []Pair{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func NumIslands(grid [][]byte) int {
	// search through grid til finding a 1
	// - perform search starting from that 1:
	//   - search recusively through all neighbour 1s
	//   - mark any neighbour 1s as 0s to show they have been visited
	// - increment counter
	// return counter

	count := 0

	numRows := len(grid)
	numCols := len(grid[0])

	var visitIsland func(node Pair)

	visitIsland = func(node Pair) {
		grid[node.Row][node.Col] = '0'

		for _, direction := range directions {
			neighbour := Pair{
				Row: node.Row + direction.Row,
				Col: node.Col + direction.Col,
			}
			if neighbour.Row < 0 ||
				neighbour.Col < 0 ||
				neighbour.Row > numRows-1 ||
				neighbour.Col > numCols-1 ||
				grid[neighbour.Row][neighbour.Col] == '0' {
				continue
			}

			visitIsland(neighbour)
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if grid[i][j] == '1' {
				pair := Pair{i, j} // Could remove pairs and just use i, j also
				visitIsland(pair)
				count++
			}
		}
	}

	return count
}
