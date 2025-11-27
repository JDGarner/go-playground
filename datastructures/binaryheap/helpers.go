package binaryheap

import (
	"fmt"
	"strings"
)

// (AI generated) String returns a pretty-printed representation of the binary heap
func (b *BinaryHeap) String() string {
	if b.Len() == 0 {
		return ""
	}
	
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%v\n", b.data))
	b.printTree(&sb, 1, "", "")
	return sb.String()
}

// (AI generated) printTree recursively builds the tree string with proper formatting
func (b *BinaryHeap) printTree(sb *strings.Builder, index int, prefix, childPrefix string) {
	// Check if index is out of bounds
	if index > b.Len() {
		return
	}
	
	// Print current node
	sb.WriteString(fmt.Sprintf("%s%d\n", prefix, b.data[index]))
	
	// Calculate child indices
	leftIndex := index * 2
	rightIndex := (index * 2) + 1
	
	// Determine if we have children
	hasLeft := leftIndex <= b.Len()
	hasRight := rightIndex <= b.Len()
	
	// Print left subtree (marked with L:)
	if hasLeft {
		if hasRight {
			// More children coming, use ├──
			b.printTree(sb, leftIndex, childPrefix+"├── L: ", childPrefix+"│      ")
		} else {
			// Last child, use └──
			b.printTree(sb, leftIndex, childPrefix+"└── L: ", childPrefix+"       ")
		}
	}
	
	// Print right subtree (marked with R:)
	if hasRight {
		// Right child is always last
		b.printTree(sb, rightIndex, childPrefix+"└── R: ", childPrefix+"       ")
	}
}
