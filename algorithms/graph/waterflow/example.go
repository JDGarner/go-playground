package waterflow

import "fmt"

func Example() {
	grid := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	// grid := [][]int{
	// 	{4, 2, 7, 3, 4},
	// 	{7, 4, 6, 4, 7},
	// 	{6, 3, 5, 3, 6},
	// }

	result := pacificAtlantic(grid)

	fmt.Println(">>> result: ", result)
}
