package adjacencylist

import (
	"fmt"
	"strings"
)

type AdjacencyList map[string][]string

// Edges example:
// [["A", "B"], ["B", "C"], ["B", "E"], ["C", "E"], ["E", "D"]]

func New(edges [][]string) AdjacencyList {
	adjList := AdjacencyList{}

	for _, edge := range edges {
		src, dst := edge[0], edge[1]

		if _, ok := adjList[src]; !ok {
			adjList[src] = []string{}
		}
		if _, ok := adjList[dst]; !ok {
			adjList[dst] = []string{}
		}

		adjList[src] = append(adjList[src], dst)
	}

	return adjList
}

// A: [B C]
// B: [A C]
// C: [E]
// E: []

// Count the number of paths from start to end, e.g. from A to E
func (a AdjacencyList) DFSCountPaths(start, end string) int {
	visited := make(map[string]struct{})

	return a.dfsCountHelper(visited, start, end)
}

func (a AdjacencyList) dfsCountHelper(visited map[string]struct{}, node, target string) int {
	if _, ok := visited[node]; ok {
		return 0
	}
	if node == target {
		return 1
	}

	count := 0

	visited[node] = struct{}{}

	for _, dst := range a[node] {
		count += a.dfsCountHelper(visited, dst, target)
	}
	delete(visited, node)

	return count
}

func (a AdjacencyList) DFS() {
	visited := make(map[string]struct{})

	for src, destinations := range a {
		a.dfsHelper(visited, src, destinations)
	}

}

func (a AdjacencyList) dfsHelper(visited map[string]struct{}, src string, dsts []string) {
	if _, ok := visited[src]; ok {
		return
	}

	visited[src] = struct{}{}

	for _, dst := range dsts {
		fmt.Printf("%s - %s\n", src, dst)
		a.dfsHelper(visited, dst, a[dst])
	}
}

func (a AdjacencyList) String() string {
	var sb strings.Builder

	for node, connections := range a {
		sb.WriteString(fmt.Sprintf("%s: %v\n", node, connections))
	}

	return sb.String()
}

// A: [B C]
// B: [A C]
// C: [E]
// E: []

// Find the length of shortest path from start to end, e.g. from A to E
func (a AdjacencyList) BFSShortestPath(start, end string) int {
	count := 0
	queue := []string{start}
	visited := map[string]struct{}{
		start: {},
	}

	for len(queue) > 0 {
		fmt.Println(">>> queue: ", queue)

		for range len(queue) {
			// Dequeue
			node := queue[0]
			queue = queue[1:]

			fmt.Println(">>> node: ", node)
			fmt.Println(">>> visited: ", visited)
			fmt.Println(">>> count: ", count)

			if node == end {
				return count
			}

			// for each neighbour, enqueue it if it has not been visited
			for _, neighbour := range a[node] {
				if _, ok := visited[neighbour]; !ok {
					queue = append(queue, neighbour)
					visited[neighbour] = struct{}{}
				}
			}
		}
		count++
	}

	return count
}
