package slidingwindow

// You are given a string s consisting of only uppercase english characters
// and an integer k.
// You can choose up to k characters of the string and replace them with
// any other uppercase English character.

// After performing at most k replacements, return the length of the longest
// substring which contains only one distinct character.

// Example 1:
// Input: s = "XYYX", k = 2
// Output: 4
// Explanation: Either replace the 'X's with 'Y's, or replace the 'Y's with 'X's.

// Example 2:
// Input: s = "AAABABB", k = 1
// Output: 5

// make k replacements in s such that we end up with the longest possible substring
// where every letter is the same

// keep track of the longest ss so far
// for each unique character keep track of it as long as there are less than k+1
// different characters in it

// in "AAABCABBBB", k=1, keep track of A until finding the BC then drop it
// keep track of first A index and number of non-A characters

// BZBXBABAAB

func characterReplacement(s string, k int) int {
	charSet := make(map[rune]struct{})

	for _, r := range s {
		if _, ok := charSet[r]; !ok {
			charSet[r] = struct{}{}
		}
	}

	max := 0

	for x := range charSet {
		left := 0
		countOfX := 0

		for right, rightChar := range s {
			if rightChar == x {
				countOfX++
			}

			// windowSize := right - left + 1
			// numOfReplacementsNeeded := windowSize - countOfX

			for ((right - left + 1) - countOfX) > k {
				if rune(s[left]) == x {
					countOfX--
				}
				left++
			}

			windowSize := right - left + 1
			if windowSize > max {
				max = windowSize
			}
		}
	}

	return max
}
