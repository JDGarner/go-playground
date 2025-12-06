package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct {
	data [][]int
}

func New(data [][]int) *Matrix {
	return &Matrix{
		data: data,
	}
}

// These are the vertices. For this example only 0s have edges between them
// {0, 0, 0, 0},
// {1, 1, 0, 0},
// {0, 0, 0, 1},
// {0, 1, 0, 0},

type Node struct {
	Row    int
	Column int
}

func (m *Matrix) isFinalNode(node Node) bool {
	return node.Row == len(m.data)-1 && node.Column == len(m.data[node.Row])-1
}

func (m *Matrix) isBlocked(node Node) bool {
	return m.data[node.Row][node.Column] != 0
}

func (m *Matrix) Get(node Node) int {
	return m.data[node.Row][node.Column]
}

func (m *Matrix) Set(node Node, value int) {
	m.data[node.Row][node.Column] = value
}

func (m *Matrix) Rows() int {
	return len(m.data)
}

func (m *Matrix) Cols() int {
	return len(m.data[0])
}

func (m *Matrix) String() string {
	var sb strings.Builder

	for _, Row := range m.data {
		sb.WriteString(fmt.Sprintf("%v\n", Row))
	}

	return sb.String()
}
