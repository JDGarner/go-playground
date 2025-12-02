package graph

import "fmt"

func MatrixDFSExample() {
	data := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	m := New(data)

	fmt.Println(m)

	uniquePaths := m.GetUniquePaths()

	fmt.Println("number of unique paths: ", len(uniquePaths))
	fmt.Println("unique paths:")
	for _, path := range uniquePaths {
		fmt.Println(path)
	}

	uniquePathCount := m.CountUniquePaths()
	fmt.Println(">>> v2 method, path count: ", uniquePathCount)


	uniquePaths = m.GetUniquePathsV2()
	fmt.Println("number of unique paths v2: ", len(uniquePaths))
	fmt.Println("unique paths v2:")
	for _, path := range uniquePaths {
		fmt.Println(path)
	}
}
