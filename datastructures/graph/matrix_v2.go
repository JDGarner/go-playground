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

func (m *Matrix) GetUniquePathsV2() [][]Node {
	start := Node{
		row:    0,
		column: 0,
	}

	visited := make(map[Node]bool)
	allPaths := [][]Node{}
	currentPath := []Node{}

	m.collectPaths(start, visited, currentPath, &allPaths)

	return allPaths
}

func (m *Matrix) collectPaths(node Node, visited map[Node]bool, currentPath []Node, allPaths *[][]Node) {
	if !m.isInBounds(node) {
		return
	}

	if m.data[node.row][node.column] != 0 {
		return
	}

	if visited[node] {
		return
	}

	currentPath = append(currentPath, node)

	if m.isFinalNode(node) {
		newPath := make([]Node, len(currentPath))
		copy(newPath, currentPath)
		*allPaths = append(*allPaths, newPath)
		return
	}

	visited[node] = true

	m.collectPaths(Node{node.row + 1, node.column}, visited, currentPath, allPaths)
	m.collectPaths(Node{node.row, node.column + 1}, visited, currentPath, allPaths)
	m.collectPaths(Node{node.row - 1, node.column}, visited, currentPath, allPaths)
	m.collectPaths(Node{node.row, node.column - 1}, visited, currentPath, allPaths)

	delete(visited, node)
}
