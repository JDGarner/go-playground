package numberboard

import "fmt"

func SnakesAndLaddersExample() {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	res := snakesAndLadders(board)
	fmt.Println(">>> res: ", res)

	board2 := [][]int{
		{-1, 4, -1},
		{6, 2, 6},
		{-1, 3, -1},
	}

	res2 := snakesAndLadders(board2)
	fmt.Println(">>> res: ", res2)
}
