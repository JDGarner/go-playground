package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

// 1: 2
// 2: 1, 3
// 3: 2

// [[2],[1,3],[2]]

// example input node:
// Val: 1
// Neighbours: [*2] (pointer to node 2)

func CloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)

	return cloneGraphHelper(node, visited)
}

func cloneGraphHelper(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if clone, ok := visited[node]; ok {
		return clone
	}

	clone := &Node{Val: node.Val}
	visited[node] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneGraphHelper(neighbor, visited))
	}

	return clone
}
