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

// type Pair struct {
// 	Row int
// 	Col int
// }

// var directions = []Pair{
// 	{0, 1},  // right
// 	{1, 0},  // down
// 	{0, -1}, // left
// 	{-1, 0}, // up
// }

// func NumIslands(grid [][]byte) int {
// 	// visited = empty map
// 	// search through grid til finding an unvisited 1
// 	// - perform search until all land visited
// 	// - increment counter

// 	// return counter

// 	visited := make(map[Pair]struct{})
// 	count := 0

// 	numRows := len(grid)
// 	numCols := len(grid[0])

// 	var visitIsland func(start Pair)

// 	visitIsland = func(start Pair) {
// 		// keep searching neighbours until no more exist
// 		queue := []Pair{start}

// 		for len(queue) > 0 {
// 			current := queue[0]
// 			queue = queue[1:]
// 			visited[current] = struct{}{}

// 			for _, direction := range directions {
// 				neighbour := Pair{
// 					Row: current.Row + direction.Row,
// 					Col: current.Col + direction.Col,
// 				}

// 				if neighbour.Row < 0 || neighbour.Col < 0 || neighbour.Row > numRows-1 || neighbour.Col > numCols-1 {
// 					continue
// 				}

// 				if grid[neighbour.Row][neighbour.Col] == '0' {
// 					continue
// 				}

// 				if _, ok := visited[neighbour]; !ok {
// 					queue = append(queue, neighbour)
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < numCols; j++ {
// 			if grid[i][j] == '0' {
// 				continue
// 			}

// 			pair := Pair{i, j}

// 			if _, ok := visited[pair]; !ok {
// 				visitIsland(pair)
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
