package matrix

// Get the unique paths to go from top left to bottom right.
// Path can only move across 0s and cannot visit the same node twice.
func (m *Matrix) GetUniquePaths() [][]Node {
	currentPath := []Node{
		{
			Row:    0,
			Column: 0,
		},
	}

	return m.getPathHelper(currentPath)
}

// {0, 0, 0, 0},
// {1, 1, 0, 0},
// {0, 0, 0, 1},
// {0, 1, 0, 0},

// Repeatedly try to take right/down/left/up paths:
//   - push onto the currentPath if the node can be taken
//   - ignore nodes already in current path / non 0 values
//
// If we are at the end - increment counter
func (m *Matrix) getPathHelper(currentPath []Node) [][]Node {
	latestNode := currentPath[len(currentPath)-1]

	if m.isFinalNode(latestNode) {
		copiedPath := make([]Node, len(currentPath))
		copy(copiedPath, currentPath)

		return [][]Node{copiedPath}
	}

	paths := [][]Node{}

	right, ok := m.takeRightNode(latestNode, currentPath)
	if ok {
		currentPath = append(currentPath, *right)
		paths = append(paths, m.getPathHelper(currentPath)...)
		currentPath = currentPath[:len(currentPath)-1]
	}

	down, ok := m.takeDownNode(latestNode, currentPath)
	if ok {
		currentPath = append(currentPath, *down)
		paths = append(paths, m.getPathHelper(currentPath)...)
		currentPath = currentPath[:len(currentPath)-1]
	}

	left, ok := m.takeLeftNode(latestNode, currentPath)
	if ok {
		currentPath = append(currentPath, *left)
		paths = append(paths, m.getPathHelper(currentPath)...)
		currentPath = currentPath[:len(currentPath)-1]
	}

	up, ok := m.takeUpNode(latestNode, currentPath)
	if ok {
		currentPath = append(currentPath, *up)
		paths = append(paths, m.getPathHelper(currentPath)...)
		currentPath = currentPath[:len(currentPath)-1]
	}

	return paths
}

func (m *Matrix) takeRightNode(node Node, currentPath []Node) (*Node, bool) {
	return m.takeNode(
		Node{
			Row:    node.Row,
			Column: node.Column + 1,
		},
		currentPath,
	)
}

func (m *Matrix) takeDownNode(node Node, currentPath []Node) (*Node, bool) {
	return m.takeNode(
		Node{
			Row:    node.Row + 1,
			Column: node.Column,
		},
		currentPath,
	)
}

func (m *Matrix) takeLeftNode(node Node, currentPath []Node) (*Node, bool) {
	return m.takeNode(
		Node{
			Row:    node.Row,
			Column: node.Column - 1,
		},
		currentPath,
	)
}

func (m *Matrix) takeUpNode(node Node, currentPath []Node) (*Node, bool) {
	return m.takeNode(
		Node{
			Row:    node.Row - 1,
			Column: node.Column,
		},
		currentPath,
	)
}

// If there is a 0 value at the given node and it is not in the currentPath return it
func (m *Matrix) takeNode(node Node, currentPath []Node) (*Node, bool) {
	if node.Row < 0 || node.Row > len(m.data)-1 {
		return nil, false
	}

	if node.Column < 0 || node.Column > len(m.data[node.Row])-1 {
		return nil, false
	}

	if m.data[node.Row][node.Column] != 0 {
		return nil, false
	}

	for _, n := range currentPath {
		if node.Row == n.Row && node.Column == n.Column {
			return nil, false
		}
	}

	return &Node{
		Row:    node.Row,
		Column: node.Column,
	}, true
}
