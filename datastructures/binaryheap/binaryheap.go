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

// Removes the root node (the one at index position 1)
// - Replace the root node with the the last node, then percolate it down:
//   - keep swapping with the minimum value until reaching a leaf node OR it is smaller than minimum
func (b *BinaryHeap) Pop() (int, bool) {
	if b.Len() == 0 {
		return 0, false
	}
	if b.Len() == 1 {
		root := b.data[1]
		b.data = b.data[:len(b.data)-1]
		return root, true
	}

	// Swap root with last element
	i := 1
	root := b.data[i]
	b.data[i] = b.data[b.Len()]

	b.data = b.data[:len(b.data)-1] // Remove last element

	for !b.IsLeafNodeIndex(i) {
		minIndex := b.GetMinIndex(leftChild(i), rightChild(i))

		if b.data[i] < b.data[minIndex] {
			return root, true
		}

		b.data[minIndex], b.data[i] = b.data[i], b.data[minIndex]
		i = minIndex
	}

	return root, true
}

func (b *BinaryHeap) Len() int {
	return len(b.data) - 1
}

func (b *BinaryHeap) GetMinIndex(left, right int) int {
	if b.Len() < right {
		return left
	}

	if b.data[left] < b.data[right] {
		return left
	}

	return right
}

func (b *BinaryHeap) IsLeafNodeIndex(index int) bool {
	return !b.HasLeftChild(index) && !b.HasRightChild(index)
}

func (b *BinaryHeap) HasLeftChild(index int) bool {
	return b.Len() >= leftChild(index)
}

func (b *BinaryHeap) HasRightChild(index int) bool {
	return b.Len() >= rightChild(index)
}

func leftChild(index int) (childIndex int) {
	return index * 2
}

func rightChild(index int) (childIndex int) {
	return (index * 2) + 1
}

func parent(index int) (parentIndex int) {
	return index / 2
}

func Heapify(data []int) *BinaryHeap {
	// Move first element to the end spot and set first index to 'null'
	data = append(data, data[0])
	data[0] = 0

	b := &BinaryHeap{
		data: data,
	}

	// start from the first node that has children (parent of last node)
	startIndex := parent(b.Len())

	// For each element in reverse
	// - Percolate it down
	for i := startIndex; i > 0; i-- {
		b.percolateDown(i)
	}

	return b
}

func (b *BinaryHeap) percolateDown(i int) {
	currentIndex := i

	for !b.IsLeafNodeIndex(currentIndex) {
		minIndex := b.GetMinIndex(leftChild(currentIndex), rightChild(currentIndex))

		if b.data[currentIndex] > b.data[minIndex] {
			b.data[currentIndex], b.data[minIndex] = b.data[minIndex], b.data[currentIndex]
			currentIndex = minIndex
		} else {
			break
		}
	}
}
