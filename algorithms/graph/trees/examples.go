package trees

import (
	"fmt"

	"github.com/JDGarner/go-playground/datastructures/binarysearchtree"
)

func KthSmallestNode() {
	tree := binarysearchtree.New(4)
	tree.SetLeft(3)
	tree.Left.SetLeft(2)
	tree.SetRight(5)

	fmt.Println(tree)

	res := kthSmallestFinal(tree, 4)
	fmt.Println(">>> res: ", res)
}
