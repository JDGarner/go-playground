package misc

// Given an integer array nums and an integer k,
// return the k most frequent elements within the array.

// The test cases are generated such that the answer is always unique.
// You may return the output in any order.

// Example 1:
// Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

// Example 2:
// Input: nums = [7,7], k = 1
// Output: [7]

func topKFrequent(nums []int, k int) []int {
	// All numbers that appear 1 time go into group freq[0].
	// All numbers that appear 2 times go into group freq[1].
	// At the end loop backwards until we have k elements

	// [7, 7]
	// freq[0] = 7
	// freq[1] = 7

	// [1,2,2,3,3,3]
	// freq[0] = 1, 2, 3
	// freq[1] = 2, 3
	// freq[2] = 3

	// To build this freq array we first need a map of nums to their count:
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	frequency := make([][]int, len(nums)+1)

	// Now build the frequency array using map
	for num, count := range counts {
		frequency[count] = append(frequency[count], num)
	}

	// Loop backwards in frequency until we collect k elements
	res := make([]int, 0, k)
	for i := len(frequency) - 1; i >= 0; i-- {
		for _, num := range frequency[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
