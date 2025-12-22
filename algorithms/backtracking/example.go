package backtracking

import "fmt"

func WordSearchExample() {
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'A', 'T'},
	// }

	// // Case 1: word is present
	// word1 := "CATS"
	// fmt.Printf("Does the word '%s' exist? %v\n", word1, exist(board, word1))

	// // Case 2: word is not present
	// word2 := "BATS"
	// fmt.Printf("Does the word '%s' exist? %v\n", word2, exist(board, word2))

	// board2 := [][]byte{
	// 	{'A', 'B', 'C', 'D'},
	// 	{'S', 'A', 'A', 'T'},
	// 	{'A', 'C', 'A', 'E'},
	// }

	// word := "CAT"
	// fmt.Printf("Does the word '%s' exist? %v\n", word, exist(board2, word))

	// board3 := [][]byte{
	// 	{'A', 'B', 'C', 'D'},
	// 	{'S', 'A', 'A', 'T'},
	// 	{'A', 'C', 'A', 'E'},
	// }

	// word = "BAT"
	// fmt.Printf("Does the word '%s' exist? %v\n", word, exist(board3, word))

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCB"
	fmt.Printf("Does the word '%s' exist? %v\n", word, exist(board, word))
	// Returns true but should return false
}
