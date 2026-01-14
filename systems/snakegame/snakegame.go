package snakegame

import "container/list"

// Design a Snake game that is played on a device with screen size height x width.
// Play the game online if you are not familiar with the game.

// The snake is initially positioned at the top left corner
// (0, 0) with a length of 1 unit.

// You are given an array food where food[i] = (rᵢ, cᵢ) is the row and column
// position of a piece of food that the snake can eat.
// When a snake eats a piece of food, its length and the game's score both increase by 1.

// Each piece of food appears one by one on the screen,
// meaning the second piece of food will not appear until the snake eats
// the first piece of food.

// When a piece of food appears on the screen, it is guaranteed that it
// will not appear on a block occupied by the snake.

// The game is over if the snake goes out of bounds (hits a wall) or if
// its head occupies a space that its body occupies after moving
// (i.e. a snake of length 4 cannot run into itself).

// Implement the SnakeGame:

// func NewSnakeGame(width, height int, food [][]int) initializes the object
// with a screen of size height x width and the positions of the food.

// func (g *SnakeGame) move(direction string) int
// returns the score of the game after applying one direction move by the snake.
// If the game is over, return -1.

// Example 1:
// Input:
// ["SnakeGame", "move", "move", "move", "move", "move", "move"]
// [[3, 2, [[1, 2], [0, 1]]], ["R"], ["D"], ["R"], ["U"], ["L"], ["U"]]

// Output:
// [null, 0, 0, 1, 1, 2, -1]

// Explanation:
// snakeGame := NewSnakeGame(3, 2, [][]int{{1, 2}, {0, 1}});
// snakeGame.move("R"); // return 0
// snakeGame.move("D"); // return 0
// snakeGame.move("R"); // return 1, snake eats the first piece of food. The second piece of food appears at (0, 1).
// snakeGame.move("U"); // return 1
// snakeGame.move("L"); // return 2, snake eats the second food. No more food appears.
// snakeGame.move("U"); // return -1, game over because snake collides with border

// Constraints:

// 1 <= width, height <= 10⁴
// 1 <= food.length <= 50
// food[i].length == 2
// 0 <= rᵢ < height
// 0 <= cᵢ < width
// direction.length == 1
// direction is 'U', 'D', 'L', or 'R'.
// At most 10⁴ calls will be made to move.

var directionMap = map[string]Cell{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

type Cell struct {
	row int
	col int
}

type SnakeGame struct {
	screen    [][]int
	food      [][]int
	snake     *Snake
	foodIndex int
	score     int
}

func NewSnakeGame(width, height int, food [][]int) *SnakeGame {
	screen := make([][]int, height)
	for row := range height {
		screen[row] = make([]int, width)
	}

	return &SnakeGame{
		screen: screen,
		food:   food,
		snake:  NewSnake(),
	}
}

func (g *SnakeGame) move(direction string) int {
	if g.score == -1 { // can't advance after game is over
		return -1
	}

	dir := directionMap[direction]
	oldHead := g.snake.head()
	newHead := Cell{
		row: oldHead.row + dir.row,
		col: oldHead.col + dir.col,
	}

	if g.isOutOfBounds(newHead) || g.snake.collidedWithSelf(newHead) {
		g.score = -1
		return -1
	}

	// if no more food left - just move the snake
	if g.foodIndex > len(g.food)-1 {
		g.snake.move(newHead)

		return g.score
	}

	currentFood := g.food[g.foodIndex]

	// if snake head is now in currentFood cell:
	// - increase score
	// - increment foodIndex
	// - grow the snake
	// else:
	// - move the snake
	if newHead.row == currentFood[0] && newHead.col == currentFood[1] {
		g.score++
		g.foodIndex++
		g.snake.grow(newHead)
	} else {
		g.snake.move(newHead)
	}

	return g.score
}

func (g *SnakeGame) isOutOfBounds(c Cell) bool {
	if c.row < 0 || c.col < 0 {
		return true
	}

	numRows, numCols := len(g.screen), len(g.screen[0])

	if c.row > numRows-1 || c.col > numCols-1 {
		return true
	}

	return false
}

// for the snake, we can count the last position as the head of the snake
// and the first as the tail. So then we can just append onto it, which is quicker than
// always inserting into the start of the slice. Could use a linked list alternatively.

type Snake struct {
	body []Cell
}

func NewSnake() *Snake {
	return &Snake{
		body: []Cell{{0, 0}},
	}
}

func (s *Snake) head() Cell {
	return s.body[len(s.body)-1]
}

func (s *Snake) move(newHead Cell) {
	// append on the new head and drop the tail (first element)
	s.body = append(s.body, newHead)
	s.body = s.body[1:]
}

func (s *Snake) grow(newHead Cell) {
	// append on the new head, no need to drop the tail
	s.body = append(s.body, newHead)
}

func (s *Snake) collidedWithSelf(newHead Cell) bool {
	for _, cell := range s.body {
		if cell.row == newHead.row && cell.col == newHead.col {
			return true
		}
	}

	return false
}

// --------------------------------------------------------------------------
// Snake version with linked list
// --------------------------------------------------------------------------

type Snake2 struct {
	body *list.List
}

func NewSnake2() *Snake2 {
	l := list.New()
	l.PushBack(Cell{0, 0})

	return &Snake2{
		body: l,
	}
}

func (s *Snake2) head() Cell {
	return s.body.Back().Value.(Cell)
}

func (s *Snake2) move(newHead Cell) {
	// append on the new head and drop the tail (first element)
	s.body.PushBack(newHead)
	s.body.Remove(s.body.Front())
}

func (s *Snake2) grow(newHead Cell) {
	// append on the new head, no need to drop the tail
	s.body.PushBack(newHead)
}

func (s *Snake2) collidedWithSelf(newHead Cell) bool {
	for e := s.body.Front(); e != nil; e = e.Next() {
		cell := e.Value.(Cell)

		if cell.row == newHead.row && cell.col == newHead.col {
			return true
		}
	}

	return false
}
