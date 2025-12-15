package waterflow

import "fmt"

func Example() {
	grid := [][]int{
		{4, 2, 7, 3, 4},
		{7, 4, 6, 4, 7},
		{6, 3, 5, 3, 6},
	}

	result := pacificAtlantic(grid)

	fmt.Println(">>> result: ", result)
}
