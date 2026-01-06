package backtracking

// You are given a string digits made up of digits from 2 through 9 inclusive.

// Each digit (not including 1) is mapped to a set of characters as shown below:

// A digit could represent any one of the characters it maps to.

// Return all possible letter combinations that digits could represent.
// You may return the answer in any order.

// Example 1:
// Input: digits = "34"
// Output: ["dg","dh","di","eg","eh","ei","fg","fh","fi"]

// for each number:
// - choose all possible char options
//   - add char to string
//   - recursively explore until we are at the last digit
//   - base case - when length of combination == len(digits)
//   - remove char from string after exploring it

// Time complexity: 
// O(4^n)
// n is the number of digits in the input string digits.
// there are 3-4 decisions at each part of the decision tree
// and the height of the tree is n (since each string will be max n length)

var digitMap = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitsRunes := []rune(digits) // if we needed to support unicode, otherwise could use bytes
	res := []string{}
	current := []rune{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(current) == len(digits) {
			res = append(res, string(current))
			return
		}

		characters := digitMap[digitsRunes[start]]

		for _, c := range characters {
			current = append(current, c)
			backtrack(start + 1) // 1
			current = current[:len(current)-1]
		}
	}

	backtrack(0)

	return res
}
