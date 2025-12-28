package array

// Given an m x n matrix, return all elements of the matrix in spiral order.

// Example 1:
// Input: matrix = [
// 	[1,2,3],
// 	[4,5,6],
// 	[7,8,9]
// ]
// Output: [1,2,3,6,9,8,7,4,5]

// Example 2:
// Input: matrix = [
// 	[1,2,3,4],
// 	[5,6,7,8],
// 	[9,10,11,12],
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

type Node struct {
	Row int
	Col int
}

var (
	right = Node{0, 1}  // right
	down  = Node{1, 0}  // down
	left  = Node{0, -1} // left
	up    = Node{-1, 0} // up

	directionToMovement = map[string]Node{
		"right": right,
		"down":  down,
		"left":  left,
		"up":    up,
	}

	nextDirection = map[string]string{
		"up":    "right",
		"right": "down",
		"down":  "left",
		"left":  "up",
	}
)

// set a right, left, top, bottom boundary
// start with direction: "right"
// switch statement on the direction
// - depending on direction, loop fully through that between boundaries (e.g. left/right boundary)
//   filling up the output array with appropriate co-ords as you go
// end when the array length = matrix size
func spiralOrder(matrix [][]int) []int {
	numRows, numCols := len(matrix), len(matrix[0])

	// initialise edge boundaries
	topBoundary, leftBoundary := 0, 0
	rightBoundary, bottomBoundary := numCols-1, numRows-1

	res := make([]int, 0, numRows*numCols)
	direction := "right"

	for len(res) < numRows*numCols {
		switch direction {
		case "right":
			for i := leftBoundary; i <= rightBoundary; i++ {
				res = append(res, matrix[topBoundary][i])
			}

			topBoundary++
		case "down":
			for i := topBoundary; i <= bottomBoundary; i++ {
				res = append(res, matrix[i][rightBoundary])
			}

			rightBoundary--
		case "left":
			for i := rightBoundary; i >= leftBoundary; i-- {
				res = append(res, matrix[bottomBoundary][i])
			}

			bottomBoundary--
		case "up":
			for i := bottomBoundary; i >= topBoundary; i-- {
				res = append(res, matrix[i][leftBoundary])
			}

			leftBoundary++
		}

		direction = nextDirection[direction]
	}

	return res
}

// set a right, left, top, bottom boundary
// start with direction: "right"
// move one step at a time filling up the output on each step
// at each step decide whether direction needs to change
func SpiralOrderFirstImpl(matrix [][]int) []int {
	numRows, numCols := len(matrix), len(matrix[0])

	// initialise edge boundaries
	topBoundary, leftBoundary := 0, -1
	rightBoundary, bottomBoundary := numCols, numRows

	current := Node{0, 0}
	res := make([]int, 0, numRows*numCols)
	res = append(res, matrix[0][0])
	direction := "right"

	for {
		// move in the current direction 1 step
		movement := directionToMovement[direction]
		next := Node{
			Row: current.Row + movement.Row,
			Col: current.Col + movement.Col,
		}

		// if we hit a boundary, change direction and increment/decrement boundary
		switch {
		case direction == "right" && next.Col == rightBoundary:
			direction = "down"
			rightBoundary--
		case direction == "down" && next.Row == bottomBoundary:
			direction = "left"
			bottomBoundary--
		case direction == "left" && next.Col == leftBoundary:
			direction = "up"
			leftBoundary++
		case direction == "up" && next.Row == topBoundary:
			direction = "right"
			topBoundary++
		default:
			res = append(res, matrix[next.Row][next.Col])
			current = next
			continue
		}

		// Need to adjust movement if direction was changed:
		movement = directionToMovement[direction]
		next = Node{
			Row: current.Row + movement.Row,
			Col: current.Col + movement.Col,
		}

		switch {
		case direction == "right" && next.Col == rightBoundary:
			return res
		case direction == "down" && next.Row == bottomBoundary:
			return res
		case direction == "left" && next.Col == leftBoundary:
			return res
		case direction == "up" && next.Row == topBoundary:
			return res
		default:
			// we haven't finished yet
		}

		current = next
		res = append(res, matrix[next.Row][next.Col])
	}
}
