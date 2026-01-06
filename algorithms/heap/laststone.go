package heap

import (
	"container/heap"
)

// You are given an array of integers stones where stones[i] represents the
// weight of the ith stone.

// We want to run a simulation on the stones as follows:

// At each step we choose the two heaviest stones, with weight x and y and
// smash them togethers
// If x == y, both stones are destroyed
// If x < y, the stone of weight x is destroyed, and the stone of weight y
// has new weight y - x.
// Continue the simulation until there is no more than one stone remaining.

// Return the weight of the last remaining stone or return 0 if none remain.

// Example 1:
// Input: stones = [2,3,6,2,4]
// Output: 1

// Explanation:
// We smash 6 and 4 and are left with a 2, so the array becomes [2,3,2,2].
// We smash 3 and 2 and are left with a 1, so the array becomes [1,2,2].
// We smash 2 and 2, so the array becomes [1].

// MaxHeap implements heap.Interface for a max-heap of ints
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	length := len(old)
	last := old[length-1]
	*h = old[0 : length-1]

	return last
}

func lastStoneWeight(stones []int) int {
	// Create and initialize max heap
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for maxHeap.Len() > 1 {
		// choose the two heaviest stones, with weight x and y and smash them togethers
		// If x == y, both stones are destroyed
		// If x < y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.

		biggest := heap.Pop(maxHeap).(int)
		secondBiggest := heap.Pop(maxHeap).(int)

		if biggest != secondBiggest {
			heap.Push(maxHeap, biggest-secondBiggest)
		}
	}

	if maxHeap.Len() == 1 {
		return maxHeap.Pop().(int)
	}

	return 0
}
