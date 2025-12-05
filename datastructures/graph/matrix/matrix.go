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
	row    int
	column int
}

func (m *Matrix) isFinalNode(node Node) bool {
	return node.row == len(m.data)-1 && node.column == len(m.data[node.row])-1
}

func (m *Matrix) isBlocked(node Node) bool {
	return m.data[node.row][node.column] != 0
}

func (m *Matrix) String() string {
	var sb strings.Builder

	for _, row := range m.data {
		sb.WriteString(fmt.Sprintf("%v\n", row))
	}

	return sb.String()
}
