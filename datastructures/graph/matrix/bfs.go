package matrix

import (
	"github.com/JDGarner/go-playground/datastructures/queue"
)

// Helper matrix:
// {0, 0, 0, 0},
// {1, 1, 0, 0},
// {0, 0, 0, 1},
// {0, 1, 0, 0},

func (m *Matrix) FindShortestPath() []Node {
	start := Node{0, 0}
	q := queue.New[Node]()
	q.Enqueue(start)

	visited := map[Node]bool{
		start: true,
	}

	// Mapping of each node to how we arrived there
	childToParent := make(map[Node]Node)

	for q.Len() > 0 {
		// Loop through whatever is on how queue at this moment
		// - which will contain the nodes that we are able to reach in 'length' steps
		//   (because we increment the length each time we have queued a particular node's neighbours)
		nodesInLayer := q.Len()

		for range nodesInLayer {
			node := q.Dequeue()

			if m.isFinalNode(node) {
				return reconstructPath(childToParent, start, node)
			}

			for _, direction := range directions {
				neighbour := Node{
					row:    node.row + direction.row,
					column: node.column + direction.column,
				}

				if !m.isInBounds(neighbour) || m.isBlocked(neighbour) || visited[neighbour] {
					continue
				}

				visited[neighbour] = true
				childToParent[neighbour] = node // <-- record how we got here
				q.Enqueue(neighbour)
			}
		}
	}

	return nil
}

func reconstructPath(childToParent map[Node]Node, start, end Node) []Node {
	current := end
	path := []Node{end}

	for current.row != start.row || current.column != start.column {
		path = append(path, childToParent[current])
		current = childToParent[current]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (m *Matrix) FindShortestPathLength() int {
	start := Node{0, 0}
	q := queue.New[Node]()
	q.Enqueue(start)

	visited := map[Node]bool{
		start: true,
	}

	length := 0

	for q.Len() > 0 {
		// Loop through whatever is on how queue at this moment
		// - which will contain the nodes that we are able to reach in 'length' steps
		//   (because we increment the length each time we have queued a particular node's neighbours)
		nodesInLayer := q.Len()

		for range nodesInLayer {
			node := q.Dequeue()

			if m.isFinalNode(node) {
				return length
			}

			for _, direction := range directions {
				neighbour := Node{
					row:    node.row + direction.row,
					column: node.column + direction.column,
				}

				if !m.isInBounds(neighbour) || m.isBlocked(neighbour) || visited[neighbour] {
					continue
				}

				q.Enqueue(neighbour)
				visited[neighbour] = true
			}
		}

		length += 1
	}

	return -1
}

var directions = []Node{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}
