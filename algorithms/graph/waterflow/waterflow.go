package waterflow

// You are given a rectangular island heights where heights[r][c] represents
// the height above sea level of the cell at coordinate (r, c).

// The islands borders the Pacific Ocean from the top and left sides,
// and borders the Atlantic Ocean from the bottom and right sides.

// Water can flow in four directions (up, down, left, or right) from a
// cell to a neighboring cell with height equal or lower. Water can also
// flow into the ocean from cells adjacent to the ocean.

// Find all cells where water can flow from that cell to both the Pacific and
// Atlantic oceans. Return it as a 2D list where each element is a list [r, c]
// representing the row and column of the cell.
// You may return the answer in any order.

// Example 1:
// Input: heights = [
//   [4,2,7,3,4],
//   [7,4,6,4,7],
//   [6,3,5,3,6]
// ]

// Output: [
// 	[0,2],[0,4],[1,0],[1,1],[1,2],[1,3],[1,4],[2,0]]

// Example 2:
// Input: heights = [[1],[1]]

// Output: [[0,0],[0,1]]

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

// - for each atlanic ocean cell (bottom & right):
//   - visit all the cells that can be reached, mark it as reached until can't
//     continue
// - repeat for pacific
// - keep track of alanticVisited and pacificVisited
// - return union of alanticVisited and pacificVisited

func pacificAtlantic(heights [][]int) [][]int {
	numRows, numCols := len(heights), len(heights[0])
	res := [][]int{}

	atlanticVisited := map[Pair]bool{}
	pacificVisited := map[Pair]bool{}

	var dfs func(p Pair, visited map[Pair]bool)
	dfs = func(p Pair, visited map[Pair]bool) {
		for _, direction := range directions {
			neighbour := Pair{p.Row + direction.Row, p.Col + direction.Col}
			if neighbour.Row < 0 || neighbour.Col < 0 {
				continue
			}
			if neighbour.Row >= numRows || neighbour.Col >= numCols {
				continue
			}
			if visited[neighbour] {
				continue
			}
			// if neighbour is more or equal than p then neighbour's water can flow down into p
			if heights[neighbour.Row][neighbour.Col] >= heights[p.Row][p.Col] {
				visited[neighbour] = true
				dfs(neighbour, visited)
			}
		}

	}

	// Bottom row (atlantic)
	for c := range numCols {
		start := Pair{numRows - 1, c}
		atlanticVisited[start] = true
		dfs(start, atlanticVisited)
	}

	// Right row (atlantic)
	for r := range numRows {
		start := Pair{r, numCols - 1}
		atlanticVisited[start] = true
		dfs(start, atlanticVisited)
	}

	// Top row (pacific)
	for c := range numCols {
		start := Pair{0, c}
		pacificVisited[start] = true
		dfs(start, pacificVisited)
	}

	// Left row (pacific)
	for r := range numRows {
		start := Pair{r, 0}
		pacificVisited[start] = true
		dfs(start, pacificVisited)
	}

	for pair := range atlanticVisited {
		if pacificVisited[pair] {
			res = append(res, []int{pair.Row, pair.Col})
		}
	}

	return res
}

// For each cell:
// - is there a path to the top or left AND bottom or right

func PacificAtlanticFirstImpl(heights [][]int) [][]int {
	result := [][]int{}
	rows, cols := len(heights), len(heights[0])

	var bfs func(node Pair) bool

	bfs = func(start Pair) bool {
		visited := make(map[Pair]bool)
		queue := []Pair{start}

		foundPacific := false
		foundAtlantic := false

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, direction := range directions {
				neighbour := Pair{
					Row: node.Row + direction.Row,
					Col: node.Col + direction.Col,
				}
				if visited[neighbour] {
					continue
				}
				if neighbour.Row == rows || neighbour.Col == cols {
					foundAtlantic = true
					if foundPacific && foundAtlantic {
						return true
					}
					continue
				}
				if neighbour.Col < 0 || neighbour.Row < 0 {
					foundPacific = true
					if foundPacific && foundAtlantic {
						return true
					}
					continue
				}
				// if cannot flow down
				if heights[node.Row][node.Col] < heights[neighbour.Row][neighbour.Col] {
					continue
				}

				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}

		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			node := Pair{i, j}

			if bfs(node) {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
