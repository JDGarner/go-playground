package courseschedule

import "fmt"

func CanFinishExample() {
	numCourses := 5
	prerequisites := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {1, 3}, {3, 4},
	}

	res := canFinish(numCourses, prerequisites)

	fmt.Println("numCourses: ", numCourses)
	fmt.Println("prerequisites: ", prerequisites)
	fmt.Println("res: ", res)
}
