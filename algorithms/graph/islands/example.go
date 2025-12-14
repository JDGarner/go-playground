package islands

import "fmt"

func Example() {
	// grid := [][]byte{
	// 	{'0', '1', '1', '1', '0'},
	// 	{'0', '1', '0', '1', '0'},
	// 	{'1', '1', '0', '0', '0'},
	// 	{'0', '0', '0', '0', '0'},
	// }
	grid := [][]byte{
		{'1', '1', '0', '0', '1'},
		{'1', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	numIslands := NumIslands(grid)

	fmt.Println(">>> numIslands: ", numIslands)

	grid2 := [][]int{
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1},
	}
	// grid2 := [][]int{
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{0, 0, 0, 1, 1},
	// 	{0, 0, 0, 1, 1},
	// }

	maxArea := MaxAreaOfIsland(grid2)

	fmt.Println(">>> maxArea: ", maxArea)
}
