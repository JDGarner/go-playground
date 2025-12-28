package numberboard

// You are given an n x n integer matrix board where the cells are labeled
// from 1 to n2 in a Boustrophedon style, e.g:
// E.g:
// 12, 11, 10, 9
// 8,  7,  6,  5
// 1,  2,  3,  4

// You start on square 1 of the board.
// In each move, do the following:
// - roll a die from 1-6, move that many spaces
// - if landing spot has a snake or ladder, follow it

// The game ends when you reach the square n2 (top left).
// if board[r][c] = -1 this means no snake or ladder
// if board[r][c] != -1, e.g. board[r][c] = 15,
// that means the snake/ladder leads to tile 15
// Squares 1 and n2 are not the starting points of any snake or ladder.

// Note that you only take a snake or ladder at most once per dice roll.
// If the destination to a snake or ladder is the start of another
// snake or ladder, you do not follow the subsequent snake or ladder.

// Return the least number of dice rolls required to reach the square n2.
// If it is not possible to reach the square, return -1.

// Example:
// Input: board = [
// 	[-1,-1,-1,-1,-1,-1],
// 	[-1,-1,-1,-1,-1,-1],
// 	[-1,-1,-1,-1,-1,-1],
// 	[-1,35,-1,-1,13,-1],
// 	[-1,-1,-1,-1,-1,-1],
// 	[-1,15,-1,-1,-1,-1]
// ]
// Output: 4

// at each decision point, could take 1-6 steps
// need to find the shortest path
// - BFS
// 'neighbours' in this case are tiles you can reach from 1-6

// fmt.Println(">>> with numRolls: ", numRolls)
// fmt.Println(">>> places reachable: ", queue)

func snakesAndLadders(board [][]int) int {
	flatBoard := flatten(board)

	numRolls := 0
	queue := []int{0}
	visited := make([]bool, len(flatBoard))

	for len(queue) > 0 {
		for range len(queue) {
			current := queue[0]
			queue = queue[1:]

			for roll := 1; roll < 7; roll++ {
				landing := getFinalLandingTile2(flatBoard, current, roll)

				if landing == len(flatBoard)-1 {
					return numRolls + 1
				}

				if visited[landing] {
					continue
				}

				visited[landing] = true
				queue = append(queue, landing)
			}
		}
		numRolls++
	}

	return -1
}

func flatten(board [][]int) []int {
	numRows, numCols := len(board), len(board[0])
	flat := make([]int, 0, numRows*numCols)

	// Loop through the rows in reverse
	for r := numRows - 1; r >= 0; r-- {

		direction := getDirection2(numRows, r)

		// depending on the direction, fill up the array going forward/backward
		if direction == 1 {
			for i := 0; i < numCols; i++ {
				flat = append(flat, board[r][i])
			}
		} else {
			for i := numCols - 1; i >= 0; i-- {
				flat = append(flat, board[r][i])
			}
		}
	}

	return flat
}

func getFinalLandingTile2(flatBoard []int, current, numSteps int) int {
	landing := current + numSteps
	boardValue := flatBoard[landing]

	// if value is a snake/ladder get the tile with that number of steps
	if boardValue != -1 {
		return boardValue - 1
	}

	return landing
}

func getDirection2(numRows, currentRow int) int {
	// if total number of rows is even, then we move right on odd rows
	if numRows%2 == 0 {
		// (e.g. on a 4x4 grid bottom row is index 3 and we move right)
		if currentRow%2 == 0 {
			return -1
		}
		return 1
	}

	// Otherwise we move right on even rows
	if currentRow%2 == 0 {
		return 1
	}
	return -1
}

type Tile struct {
	Row int
	Col int
}

func snakesAndLaddersFirstImpl(board [][]int) int {
	numRows, numCols := len(board), len(board[0])
	numRolls := 0
	queue := []Tile{
		{numRows - 1, 0},
	}
	visited := make([][]bool, len(board))
	for r := range numRows {
		visited[r] = make([]bool, numCols)
	}

	finalTile := Tile{0, 0}
	// if odd number of rows, final tile will be in the top right
	if numRows%2 != 0 {
		finalTile = Tile{0, numCols - 1}
	}

	for len(queue) > 0 {
		for range len(queue) {
			current := queue[0]
			queue = queue[1:]

			for roll := 1; roll < 7; roll++ {
				landing := getFinalLandingTile(board, current, roll)

				if landing.Row == finalTile.Row && landing.Col == finalTile.Col {
					return numRolls + 1
				}

				if visited[landing.Row][landing.Col] {
					continue
				}

				visited[landing.Row][landing.Col] = true
				queue = append(queue, landing)
			}
		}
		numRolls++
	}

	return -1
}

func getFinalLandingTile(board [][]int, current Tile, numSteps int) Tile {
	tile := getNthTileFromCurrent(board, current, numSteps)
	boardValue := board[tile.Row][tile.Col]

	// if value is a snake/ladder get the tile with that number of steps
	if boardValue != -1 {
		return getNthTileFromCurrent(board, Tile{len(board) - 1, 0}, boardValue-1)
	}

	return tile
}

func getNthTileFromCurrent(board [][]int, current Tile, n int) Tile {
	for range n {
		direction := getDirection(len(board), current)

		current = getSingleStep(board, current, direction)
	}

	return current
}

func getSingleStep(board [][]int, current Tile, direction int) Tile {
	numRows, numCols := len(board), len(board[0])

	finalTile := Tile{0, 0}
	// if odd number of rows, final tile will be in the top right
	if numRows%2 != 0 {
		finalTile = Tile{0, numCols - 1}
	}

	// If we are at the end, return end (so we don't go out of bounds)
	if current.Col == finalTile.Col && current.Row == finalTile.Row {
		return finalTile
	}

	// If we are on the last column going right, we need to go up
	if current.Col == numCols-1 && direction == 1 {
		return Tile{
			Row: current.Row - 1,
			Col: current.Col,
		}
	}

	// If we are on the first column going left, we need to go up
	if current.Col == 0 && direction == -1 {
		return Tile{
			Row: current.Row - 1,
			Col: current.Col,
		}
	}

	// Move either 1 to the left or right depending on direction
	return Tile{
		Row: current.Row,
		Col: current.Col + direction,
	}
}

func getDirection(numRows int, current Tile) int {
	// if total number of rows is even, then we move right on odd rows
	if numRows%2 == 0 {
		// (e.g. on a 4x4 grid bottom row is index 3 and we move right)
		if current.Row%2 == 0 {
			return -1
		}
		return 1
	}

	// Otherwise we move right on even rows
	if current.Row%2 == 0 {
		return 1
	}
	return -1
}
