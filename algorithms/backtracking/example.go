package backtracking

import "fmt"

func WordSearchExample() {
	board := [][]byte{
		{'A', 'B', 'C', 'D'},
		{'S', 'A', 'A', 'T'},
		{'A', 'C', 'A', 'E'},
	}

	word := "BAT"
	fmt.Printf("Does the word '%s' exist? %v\n", word, exist(board, word))

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word = "ABCB"
	fmt.Printf("Does the word '%s' exist? %v\n", word, exist(board, word))
}

func WordSearch2Example() {
	board := [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'},
	}

	words := []string{"oa", "oaa"}
	fmt.Printf("Which words from '%v' exist? %v\n", words, findWords(board, words))
}
