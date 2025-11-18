package binarysearchtree

import (
	"fmt"
	"strings"
)

// (AI generated) String returns a pretty-printed representation of the binary search tree
func (b *BST) String() string {
	if b == nil {
		return ""
	}

	var sb strings.Builder
	b.printTree(&sb, "", "")
	return sb.String()
}

// (AI generated) printTree recursively builds the tree string with proper formatting
func (b *BST) printTree(sb *strings.Builder, prefix, childPrefix string) {
	if b == nil {
		return
	}

	// Print current node
	sb.WriteString(fmt.Sprintf("%s%d\n", prefix, b.Value))

	// Determine if we have children
	hasLeft := b.Left != nil
	hasRight := b.Right != nil

	// Print left subtree (marked with L:)
	if hasLeft {
		if hasRight {
			// More children coming, use ├──
			b.Left.printTree(sb, childPrefix+"├── L: ", childPrefix+"│      ")
		} else {
			// Last child, use └──
			b.Left.printTree(sb, childPrefix+"└── L: ", childPrefix+"       ")
		}
	}

	// Print right subtree (marked with R:)
	if hasRight {
		// Right child is always last
		b.Right.printTree(sb, childPrefix+"└── R: ", childPrefix+"       ")
	}
}
