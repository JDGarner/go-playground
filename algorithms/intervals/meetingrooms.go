package intervals

import (
	"slices"
	"sort"
)

// Given an array of meeting time interval objects consisting of
// start and end times [[start_1,end_1],[start_2,end_2],...] (start_i < end_i),
// find the minimum number of days required to schedule all
// meetings without any conflicts.

// Note: (0,8),(8,10) is not considered a conflict at 8.

// Example 1:
// Input: intervals = [(0,40),(5,10),(15,20)]
// Output: 2

// Explanation:
// day1: (0,40)
// day2: (5,10),(15,20)

// Example 2:
// Input: intervals = [(4,9)]
// Output: 1

// You should aim for a solution with O(nlogn) time and O(n) space,
// where n is the size of the input array.

// Input: intervals = [(15,20),(0,40),(5,10)]

// start putting them on schedule[]
// - would need to loop through schedule each time to find where to put it
// - a the mid point we would be looping through n/2 for each n - is this O(n^2)

// sort the intervals first by start time
// [(15,20),(0,40),(5,10)]
// [(0,40),(5,10),(15,20)]

// could then schedule one by one

// Try to visualize the meetings as line segments on a number line representing
// start and end times.
// The number of rooms required is the maximum number of overlapping meetings
// at any point on the number line.
// Can you think of a way to determine this efficiently?

// ------------------------------
// 0                            40
//     5  10
//            15    20
//                  20 22

// 0   5     15     20
//        10        20 22       40

// max = 0
// concurrent = 0

// startTimes = [0, 5, 15, 20]
// endTime =    [10, 20, 22, 40]

// Keep track of the concurrent meetings iterating through the number line
// if a meeting starts before another has ended => concurrent++
// e.g. 0 vs 10 => 0 has started before 10 as ended => concurrent++
// e.g. 5 vs 10 => 5 has started before 10 as ended => concurrent++
// e.g. 15 vs 10 => 15 has not started before 10 as ended => concurrent++

type Interval struct {
	start int
	end   int
}

func minMeetingRooms(intervals []Interval) int {
	startTimes := make([]int, 0, len(intervals))
	endTimes := make([]int, 0, len(intervals))

	for _, interval := range intervals {
		startTimes = append(startTimes, interval.start)
		endTimes = append(endTimes, interval.end)
	}

	slices.Sort(startTimes)
	slices.Sort(endTimes)

	startPointer := 0
	endPointer := 0
	max := 0        // maximum of concurrent during the whole loop
	concurrent := 0 // how many meetings are currently going on at the same time

	// whenever a meeting starts before another one ends, we need a new room
	// whenever a meeting ends before or at the same time another starts, a room becomes free

	for startPointer < len(intervals) {
		if startTimes[startPointer] < endTimes[endPointer] {
			concurrent++
			startPointer++
		} else {
			concurrent--
			endPointer++
		}

		if concurrent > max {
			max = concurrent
		}
	}

	return max
}

/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms2(intervals []Interval) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))

	for i, interval := range intervals {
		start[i] = interval.start
		end[i] = interval.end
	}

	sort.Ints(start)
	sort.Ints(end)

	res, count := 0, 0
	s, e := 0, 0

	for s < len(intervals) {
		if start[s] < end[e] {
			s++
			count++
		} else {
			e++
			count--
		}
		if count > res {
			res = count
		}
	}

	return res
}
