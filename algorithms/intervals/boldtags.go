package intervals

import (
	"sort"
	"strings"
)

// You are given a string s and an array of strings words.

// You should add a closed pair of bold tag <b> and </b> to wrap the substrings in s that exist in words.

// If two such substrings overlap, you should wrap them together with only one pair of closed bold-tag.
// If two substrings wrapped by bold tags are consecutive, you should combine them.
// Return s after adding the bold tags.

// Example 1:

// Input: s = "abcxyz123", words = ["abc","123"]

// Output: "<b>abc</b>xyz<b>123</b>"
// Explanation: The two strings of words are substrings of s as following: "abcxyz123".
// We add before each substring and after each substring.

// Example 2:

// Input: s = "aaabbb", words = ["aa","b"]

// Output: "<b>aaabbb</b>"
// Explanation:
// "aa" appears as a substring two times: "aaabbb" and "aaabbb".
// "b" appears as a substring three times: "aaabbb", "aaabbb", and "aaabbb".
// Since the first two 's overlap, we merge them: "aaaabbb".
// Since now the four 's are consecutive, we merge them: "aaabbb".

// loop through each word and get all the intervals that it appears
// merge those intervals together into one intervals array e.g. [[1, 3], [7, 10]]
// insert the b tags into each interval

const (
	bTagOpen  = "<b>"
	bTagClose = "</b>"
)

func AddBoldTag(s string, words []string) string {
	intervals := [][]int{}

	for _, word := range words {
		// for i := 0; i < len(s); i++ {
		// 	if strings.HasPrefix(s[i:], word) {
		// 		start, end := i, i+len(word)
		// 		intervals = append(intervals, []int{start, end})
		// 	}
		// }

		idx := 0
		for {
			pos := strings.Index(s[idx:], word)
			if pos == -1 {
				break
			}
			start := idx + pos
			end := start + len(word)
			intervals = append(intervals, []int{start, end})
			idx = start + 1
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := merge(intervals)

	// insert the b tags into each interval
	var sb strings.Builder
	i := 0

	for _, interval := range merged {
		before := string(s[i:interval[0]])
		sb.WriteString(before)
		sb.WriteString(bTagOpen)

		i = interval[0]

		after := string(s[i:interval[1]])
		sb.WriteString(after)
		sb.WriteString(bTagClose)

		i = interval[1]
	}

	sb.WriteString(string(s[i:]))

	return sb.String()
}
