package twopointers

import "sort"

// You are given an integer array people where people[i] is the weight of the
// ith person, and an infinite number of boats where each boat can carry
// a maximum weight of limit. Each boat carries at most two people at the same time,
// provided the sum of the weight of those people is at most limit.

// Return the minimum number of boats to carry every given person.

// Example 1:
// Input: people = [5,1,4,2], limit = 6

// Output: 2
// Explanation:
// First boat [5,1].
// Second boat [4,2].

// Example 2:

// Input: people = [1,3,2,3,2], limit = 3

// Output: 4
// Explanation:
// First boat [3].
// Second boat [3].
// Third boat [1,2].
// Fourth boat [2].

func numRescueBoats(people []int, limit int) int {
	// ideally each boat carries as close to limit as possible (heaviest+lightest)

	// sort the people
	// increment in from the right and left, putting people in boat

	// sort from lightest to heaviest
	sort.Slice(people, func(i, j int) bool {
		return people[i] < people[j]
	})

	boats := 0
	left, right := 0, len(people)-1

	// [5,1,4,2]

	for left < right {
		// they both fit in
		if people[left]+people[right] <= limit {
			left++
			right--
		} else {
			right-- // only the heavier one fits in
		}

		boats++
	}

	// at this point either left == right, in which case there is one more to put in
	if left == right {
		return boats + 1
	}

	// or left > right, which means we put everyone in already
	return boats
}

func numRescueBoats2(people []int, limit int) int {
	// Find the maximum weight to determine counting array size
	maxWeight := 0
	for _, weight := range people {
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	// Counting sort: count frequency of each weight
	count := make([]int, maxWeight+1)
	for _, weight := range people {
		count[weight]++
	}

	// e.g. for  [5,1,4,5,5]
	// would be: [0,1,0,0,1,3] (e.g. 3 with weight 5)

	// Two-pointer approach on the counting array
	boats := 0
	left := 1          // Start from weight 1
	right := maxWeight // Start from max weight

	for left <= right {
		// Find the next lightest person
		for left <= right && count[left] == 0 {
			left++
		}
		// Find the next heaviest person
		for left <= right && count[right] == 0 {
			right--
		}

		if left > right {
			break
		}

		// Try to pair lightest with heaviest
		if left+right <= limit {
			count[left]--
			count[right]--
			boats++

			// If both pointers point to same weight and we just used both
			if left == right && count[left] == 0 {
				left++
				right--
			}
		} else {
			// Only heaviest person goes in boat
			count[right]--
			boats++
		}
	}

	return boats
}
