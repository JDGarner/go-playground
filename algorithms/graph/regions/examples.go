package regions

import (
	"fmt"
)

func SurroundedRegionsExample() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	fmt.Println("Before:")
	for _, row := range board {
		fmt.Println(string(row))
	}

	solve(board)

	fmt.Println("\nAfter:")
	for _, row := range board {
		fmt.Println(string(row))
	}

	// board := [][]byte{
	// 	{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	// 	{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	// 	{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
	// 	{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
	// 	{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	// }

	// expected := [][]byte{
	// 	{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	// 	{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
	// 	{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
	// 	{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
	// 	{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
	// 	{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	// }

	// fmt.Println("Before:")
	// for _, row := range board {
	// 	fmt.Println(string(row))
	// }

	// solve(board)

	// fmt.Println("\nAfter:")
	// for _, row := range board {
	// 	fmt.Println(string(row))
	// }
}
