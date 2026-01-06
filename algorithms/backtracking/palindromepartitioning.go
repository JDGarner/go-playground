package backtracking

import (
	"fmt"
)

// Given a string s, split s into substrings where every substring
// is a palindrome.
// Return all possible lists of palindromic substrings.
// You may return the solution in any order.

// Example 1:
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]

// Example 2:
// Input: s = "a"
// Output: [["a"]]

func partition2(s string) [][]string {
	// res stores all valid palindrome partitions
	res := [][]string{}
	// part stores the current partition being built
	part := []string{}

	// dfs explores all possible partitions using two pointers:
	// j = start index of current substring being considered
	// i = end index of current substring being considered
	var dfs func(start, end int)
	dfs = func(start, end int) {
		// Base case: we've gone past the end of the string
		if end >= len(s) {
			// Check if we've consumed all characters (no leftover substring)
			// i == j means we successfully partitioned the entire string
			if end == start {
				// Create a copy of part and add to results
				// (copying is necessary because part is modified during backtracking)
				res = append(res, append([]string{}, part...))
			}
			return
		}

		// "a"

		// part = ["a"] then look from 1,1
		// then undo and remove the part

		// search from 

		// CHOICE 1: Try to partition here if s[start:end+1] is a palindrome
		if isPali(s, start, end) {
			// Add the palindromic substring to current partition
			part = append(part, s[start:end+1])

			// Recurse to find partitions for the rest of the string
			// Start the next substring at end+1 (because we've already added in all the way to end)
			dfs(end+1, end+1)

			// Backtrack: remove the substring we just added
			// to explore other partitioning possibilities
			part = part[:len(part)-1]
		}

		// CHOICE 2: Don't partition here, extend the current substring
		// Move the end pointer forward to consider a longer substring
		dfs(start, end+1)
	}

	// Start DFS with both pointers at the beginning
	dfs(0, 0)
	return res
}

// isPali checks if s[l:r+1] is a palindrome using two pointers
func isPali(s string, l, r int) bool {
	// Compare characters from both ends moving inward
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// Explore all partitions

// at each point you can either start a partition OR add onto current string
// 2 decision options
// => decision will be made n times (where n = length of string)
// so possibilities = 2^n
// e.g. for length 3 string 2^3 = 8 possibilities

// Explore all partition possiblities using backtracking
// - only decide to partition if current string is a palindrome
//   - otherwise continue
//   - if at the end => if latest is a palindrome => add the set of partitions

func partition(s string) [][]string {
	res := [][]string{}
	current := []string{}

	sRunes := []rune(s)

	var backtrack func(start, wordIndex int)
	backtrack = func(start, wordIndex int) {
		if start == len(sRunes) {
			if !isPalindrome(current[len(current)-1]) {
				return
			}

			temp := append([]string{}, current...)
			res = append(res, temp)
			return
		}

		// Input: s = "aba"
		// current = ["a", ]
		// backtrack(1, 1)

		// CHOICE 1 - WHETHER TO PARTITION
		// only decide to partition if current string is a palindrome
		if len(current) > 0 && isPalindrome(current[wordIndex]) {
			current = append(current, string(s[start]))
			backtrack(start+1, wordIndex+1)

			// todo: do backtrack undo step - remove latest partition
			current = current[:len(current)-1]
		}

		// CHOICE 2 - WHETHER TO NOT PARTITION (ADD TO CURRENT STRING)
		// always add to current string because it may become a palindrome
		if len(current) <= wordIndex {
			current = append(current, "")
		}
		current[wordIndex] = fmt.Sprintf("%s%s", current[wordIndex], string(s[start]))
		backtrack(start+1, wordIndex)

		// todo: do backtrack undo step - remove latest character from current string
		current[wordIndex] = current[wordIndex][:len(current[wordIndex])-1]
	}

	backtrack(0, 0)

	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
