package dynamic

import "fmt"

// Given two strings text1 and text2, return the length of the longest common
// subsequence between the two strings if one exists, otherwise return 0.

// A subsequence is a sequence that can be derived from the given sequence
// by deleting some or no elements without changing the relative order of
// the remaining characters.

// For example, "cat" is a subsequence of "crabt".
// A common subsequence of two strings is a subsequence that exists in both strings.

// Example 1:
// Input: text1 = "cat", text2 = "crabt"
// Output: 3

// Explanation: The longest common subsequence is "cat" which has a length of 3.

// Example 2:
// Input: text1 = "abcd", text2 = "abcd"
// Output: 4

// Example 3:
// Input: text1 = "abcd", text2 = "efgh"
// Output: 0

// cat
// crabt

// zcat
// crabt

//              zc
//     zr             cc
//  za   rc         cr. ca

//           z              c
//.  look for next z

// zc
// c

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// memo[i][j] stores the length of the longest common subsequence
	// starting from position i in text1 and position j in text2.
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			if text1[i] == text2[j] {
				memo[i][j] = 1 + memo[i+1][j+1]
			} else {
				memo[i][j] = max(memo[i][j+1], memo[i+1][j])
			}
		}
	}

	for i := range memo {
		fmt.Println(memo[i])
	}

	return memo[0][0]
}

func longestCommonSubsequenceTopDown(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// memo[i][j] stores the length of the longest common subsequence
	// starting from position i in text1 and position j in text2.
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == m || j == n {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// case 1:
		// characters match
		// - advance both counters, add 1 for the match
		if text1[i] == text2[j] {
			memo[i][j] = dfs(i+1, j+1) + 1
		} else {
			memo[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}

		return memo[i][j]
	}

	return dfs(0, 0)
}
