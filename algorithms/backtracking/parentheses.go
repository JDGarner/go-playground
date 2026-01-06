package backtracking

import (
	"fmt"
	"strings"
)

// You are given an integer n.
// Return all well-formed parentheses strings that you can generate with
// n pairs of parentheses.

// Example 1:
// Input: n = 1
// Output: ["()"]

// Example 2:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// You may return the answer in any order.

// each n adds another open and close in decision tree

// at each point, do you either open or close?
// if you open a new one, have to close later on

// Input: n = 2
// Output: ["(())", "()()"]

// keep track of cursor index within the string

// cursor = 1
// (|) from this point can do:

// ((|)) <- option 1 insert before and after cursor, cursor++
// ()|() <- option 2 move cursor 1 and insert in front
// (()|) <- option 3 insert behind cursor, cursor+=2

// (()|())
// base case = len currentString == n*2

// when you open => add (

// n = 2
//       (
//  (       )

//

func generateParenthesis(n int) []string {
	stack := make([]string, 0)
	res := make([]string, 0)

	var backtrack func(int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			res = append(res, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}

		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return res
}

func generateParenthesis1(n int) []string {
	res := []string{}
	current := []rune{}

	var backtrack func()
	backtrack = func() {
		if len(current) == n*2 {
			fmt.Println(">>> possibility: ", string(current))

			if !isValid(current) {
				return
			}

			res = append(res, string(current))
			return
		}

		// explore ALL possibilities - invalid ones will be rejected later
		current = append(current, '(')
		backtrack()
		current = current[:len(current)-1]
		current = append(current, ')')
		backtrack()
		current = current[:len(current)-1]
	}

	backtrack()

	return res
}

func isValid(v []rune) bool {
	stackCount := 0

	for _, c := range v {
		if c == '(' {
			stackCount++
		} else {
			if stackCount == 0 {
				return false
			}
			stackCount--
		}
	}

	return stackCount == 0
}

func isValidBrackets(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
