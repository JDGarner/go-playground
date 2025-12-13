package courseschedule

// You are given an array prerequisites where prerequisites[i] = [a, b]
// indicates that you must take course b first if you want to take course a.

// The pair [0, 1], indicates that must take course 1 before taking course 0.

// There are a total of numCourses courses you are required to take,
// labeled from 0 to numCourses - 1.

// Return true if it is possible to finish all courses, otherwise return false.

// Input: numCourses = 2, prerequisites = [[0,1]]
// Output: true

// adjList example:
// 0: [1, 2]       // [1, 0], [2, 0]
// 1: []
// 2: [3]          // [2, 3]
// 3: [1]          // [3, 1]

// is there a cycle in the graph?
// - if yes - return false, else return true

// numCourses=5
// prerequisites=[[1,4],[2,4],[3,1],[3,2]]
// 1: [4]
// 2: [4]
// 3: [1, 2]
// 4: []

// numCourses=1
// prerequisites=[]

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)
	for i := range numCourses {
		adjList[i] = []int{}
	}

	for _, pair := range prerequisites {
		crs := pair[0]
		pre := pair[1]
		adjList[crs] = append(adjList[crs], pre)
	}

	visited := make(map[int]struct{})

	var dfs func(crs int) bool

	dfs = func(crs int) bool {
		if _, ok := visited[crs]; ok {
			return false
		}
		if len(adjList[crs]) == 0 {
			return true
		}
		visited[crs] = struct{}{}

		for _, pre := range adjList[crs] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visited, crs)
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
