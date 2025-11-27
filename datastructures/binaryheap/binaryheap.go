package binaryheap

// A binary heap is a tree structure for which the following is true:
// - For every node, all child nodes are higher than it
// - It is complete (no level of the tree has gaps) apart from the last level, which
//   must be filled left to right

// Example:
//         7
//    13         10
// 27    30   11

// It can be represented using an array, e.g:
// [0, 7, 13, 10, 27, 30, 11]

// Where the following is true:
// - Element at index 0 is unused
// - To get left child do: index * 2
// - To get right child do: (index * 2) + 1
// - To get parent do: Math.floor(index / 2)

type BinaryHeap struct {
	data []int
}

func New() *BinaryHeap {
	return &BinaryHeap{
		data: []int{0}, // Initialise with 'null' element at index 0
	}
}

// Add new value at the end of b.data (the first empty leaf node)
// If value is more than parent => we are finished
// Else swap with parent until either it is more than parent or we are at index 1
func (b *BinaryHeap) Push(value int) {
	b.data = append(b.data, value)

	valueIndex := len(b.data) - 1
	parentIndex := valueIndex / 2
	for parentIndex >= 1 && value < b.data[parentIndex] {
		b.data[parentIndex], b.data[valueIndex] = b.data[valueIndex], b.data[parentIndex]
		valueIndex = parentIndex
		parentIndex = valueIndex / 2
	}
}

func (b *BinaryHeap) Pop() {

}

func (b *BinaryHeap) Len() int {
	return len(b.data) - 1
}
