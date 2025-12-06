package rottenfruit

import (
	"github.com/JDGarner/go-playground/datastructures/graph/matrix"
	"github.com/JDGarner/go-playground/datastructures/queue"
)

// You are given a 2-D matrix grid. Each cell can have one of three possible values:

// 0 representing an empty cell
// 1 representing a fresh fruit
// 2 representing a rotten fruit
// Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit,
// then the fresh fruit also becomes rotten.

// Return the minimum number of minutes that must elapse until there are zero fresh fruits
// remaining. If this state is impossible within the grid, return -1.

// Example 1:
// Input: grid = [
//   [1,1,0],
//   [0,1,1],
//   [0,1,2]
// ]
// Output: 4

// Example 2:
// Input: grid = [
//   [1,0,1],
//   [0,2,0],
//   [1,0,1]
// ]
// Output: -1

// Constraints:
// 1 <= grid.length, grid[i].length <= 10

type Fruit int

const (
	Empty       Fruit = iota // 0
	FreshFruit               // 1
	RottenFruit              // 2
)

var directions = []matrix.Node{
	{Row: 0, Column: 1},  // right
	{Row: 1, Column: 0},  // down
	{Row: 0, Column: -1}, // left
	{Row: -1, Column: 0}, // up
}

// Starting from all the 2s and only being allowed to go through 1s,
// how many steps would it take to visit every 1

func RunRottenFruit(m *matrix.Matrix) int {
	q := queueAllRottenFruit(m)
	minutes := 0

	for q.Len() > 0 {
		for range q.Len() {

			rottenFruit := q.Dequeue()

			for _, direction := range directions {
				neighbour := matrix.Node{
					Row:    rottenFruit.Row + direction.Row,
					Column: rottenFruit.Column + direction.Column,
				}

				// if out of bounds, not 1 value (not a fresh fruit), already visited => continue
				if neighbour.Row < 0 || neighbour.Column < 0 ||
					neighbour.Row > m.Rows()-1 || neighbour.Column > m.Cols()-1 ||
					m.Get(neighbour) != int(FreshFruit) {
					continue
				}

				m.Set(neighbour, int(RottenFruit))
				q.Enqueue(neighbour)
			}
		}
		minutes++
	}

	if hasFreshFruit(m) {
		return -1
	}

	return minutes - 1
}

func queueAllRottenFruit(m *matrix.Matrix) *queue.Queue[matrix.Node] {
	q := queue.New[matrix.Node]()

	for row := range m.Rows() {
		for col := range m.Cols() {
			node := matrix.Node{Row: row, Column: col}
			val := m.Get(node)
			if val == int(RottenFruit) {
				q.Enqueue(node)
			}
		}
	}

	return q
}

func hasFreshFruit(m *matrix.Matrix) bool {
	for row := range m.Rows() {
		for col := range m.Cols() {
			val := m.Get(matrix.Node{Row: row, Column: col})
			if val == int(FreshFruit) {
				return true
			}
		}
	}

	return false
}
