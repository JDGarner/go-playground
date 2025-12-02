package graph

// Count the number of unique paths to go from top left to bottom right.
// Path can only move across 0s and cannot visit the same node twice.
func (m *Matrix) CountUniquePaths() int {
	start := Node{
		row:    0,
		column: 0,
	}

	visited := make(map[Node]bool)

	return m.countHelper(start, visited)
}

func (m *Matrix) countHelper(node Node, visited map[Node]bool) int {
	if !m.isInBounds(node) {
		return 0
	}

	if m.data[node.row][node.column] != 0 {
		return 0
	}

	if visited[node] {
		return 0
	}

	if m.isFinalNode(node) {
		return 1
	}

	visited[node] = true
	count := 0

	count += m.countHelper(Node{node.row + 1, node.column}, visited)
	count += m.countHelper(Node{node.row, node.column + 1}, visited)
	count += m.countHelper(Node{node.row - 1, node.column}, visited)
	count += m.countHelper(Node{node.row, node.column - 1}, visited)

	delete(visited, node)

	return count
}

func (m *Matrix) isInBounds(node Node) bool {
	if node.row < 0 || node.column < 0 {
		return false
	}

	if node.row > len(m.data)-1 || node.column > len(m.data[node.row])-1 {
		return false
	}

	return true
}
