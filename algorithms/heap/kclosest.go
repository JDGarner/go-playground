package heap

import (
	"container/heap"
	"math"
)

// You are given an 2-D array points where points[i] = [xi, yi] represents
// the coordinates of a point on an X-Y axis plane.
// You are also given an integer k.

// Return the k closest points to the origin (0, 0).

// The distance between two points is defined as the Euclidean
// distance (sqrt((x1 - x2)^2 + (y1 - y2)^2)).

// You may return the answer in any order.

// Example:
// Input: points = [[0,2],[2,0],[2,2]], k = 2
// Output: [[0,2],[2,0]]

type Point struct {
	value    []int
	distance float64
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Pop() any {
	old := *h
	length := len(old)
	last := old[length-1]
	*h = old[:length-1]

	return last
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func kClosest(points [][]int, k int) [][]int {
	// create a min heap where it is minimising based on distance from origin

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, point := range points {
		x := math.Pow(float64(point[0]), 2)
		y := math.Pow(float64(point[1]), 2)
		distance := math.Sqrt(x + y)
		
		heap.Push(minHeap, Point{
			value:    point,
			distance: distance,
		})
	}

	res := make([][]int, k)
	for i := range k {
		res[i] = heap.Pop(minHeap).(Point).value
	}

	return res
}
