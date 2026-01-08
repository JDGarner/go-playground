package binarysearch

import (
	"math"
)

// You are given an integer array piles where piles[i] is the number of bananas
// in the ith pile.
// You are also given an integer h, which represents the number of hours
// you have to eat all the bananas.

// You may decide your bananas-per-hour eating rate of k.
// Each hour, you may choose a pile of bananas and eat k bananas from that pile.
// If the pile has less than k bananas, you may finish eating the pile but you
// can not eat from another pile in the same hour.

// Return the minimum integer k such that you can eat all the bananas within h hours.

// Example 1:
// Input: piles = [1,4,3,2], h = 9
//                [1,2,2,1]
// Output: 2

// total = 10
// [1,2,3,4]
//

// Explanation:
// With an eating rate of 2, you can eat the bananas in 6 hours.
// With an eating rate of 1, you would need 10 hours to eat all the
// bananas (which exceeds h=9), thus the minimum eating rate is 2.

// Example 2:
// Input: piles = [25,10,23,4], h = 4
// Output: 25

func minEatingSpeed(piles []int, h int) int {
	right := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > right {
			right = piles[i]
		}
	}

	left := 1
	res := right

	for left <= right {
		totalTime := 0
		rate := left + (right-left)/2

		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(rate)))
		}

		// if it's possible with mid
		if totalTime <= h {
			// try again with the left half to see if we can get any better
			res = rate
			right = rate - 1

		} else {
			// try again with the right half since we can't do it with midpoint
			left = rate + 1
		}
	}

	return res
}

func minEatingSpeedFirstImpl(piles []int, h int) int {
	max := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}

	// max is the absolute maxiumum the answer (k) can be
	// because worst case we have to eat through the max pile in one hour
	// because we can only eat one pile at a time
	// e.g. [1, 1, 1, 100] if h = 3 => not possible, if h = 4, could do it in 100 per hour

	// All the possible answers:
	possibilities := make([]int, max) // from 1 to max
	for i := range max {
		possibilities[i] = i + 1
	}

	// brute force would be trying every value from 1 until max

	// could do binary search and go down if it's possible, up if not

	left, right := 0, len(possibilities)-1
	var prevCorrect int

	for left <= right {
		mid := left + (right-left)/2

		// if it's possible with mid
		if canFinishAtRate(piles, h, possibilities[mid]) {
			// try again with the left half to see if we can get any better
			prevCorrect = mid
			right = mid - 1

		} else {
			// try again with the right half since we can't do it with midpoint
			left = mid + 1
		}
	}

	return possibilities[prevCorrect]
}

func canFinishAtRate(piles []int, h, rate int) bool {
	totalTime := 0

	// 5 bananas
	// eating 2 per hour
	// 2.5 hours

	for _, pile := range piles {
		totalTime += int(math.Ceil(float64(pile) / float64(rate)))
	}

	return totalTime <= h
}
