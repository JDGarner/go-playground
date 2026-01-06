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

func CombinationSumExample() {
	res := combinationSum([]int{2, 5, 6, 9}, 9)

	fmt.Println(">>> res: ", res)
}

func BacktrackingExample() {
	// res := subsets([]int{1, 2, 3})
	// res := subsetsWithDup([]int{1, 2, 1})

	// fmt.Println(">>> res: ", res)

	// res := permute([]int{1, 2, 3})
	// fmt.Println(">>> res: ", res)

	// res := combine(3, 2)
	// fmt.Println(">>> res: ", res)

	// res := letterCombinations("34")
	// fmt.Println(">>> res: ", res)

	// res := generateParenthesis(3)
	// fmt.Println(">>> res: ", res)

	res := partition("aab")
	fmt.Println(">>> res: ", res)
}
