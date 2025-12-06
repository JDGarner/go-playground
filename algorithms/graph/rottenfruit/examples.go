package rottenfruit

import (
	"fmt"

	"github.com/JDGarner/go-playground/datastructures/graph/matrix"
)

func RottenFruitExample() {
	input := matrix.New([][]int{
		{1, 1, 0},
		{0, 1, 1},
		{0, 1, 2},
	})

	output := RunRottenFruit(input)

	fmt.Println("input:")
	fmt.Println(input)
	fmt.Println("output:")
	fmt.Println(output)

	// input = matrix.New([][]int{
	// 	{1, 0, 1},
	// 	{0, 2, 0},
	// 	{1, 0, 1},
	// })

	// output = RunRottenFruit(input)

	// fmt.Println("input:")
	// fmt.Println(input)
	// fmt.Println("output:")
	// fmt.Println(output)
}
