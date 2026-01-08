package dynamic

// You are given an integer array coins representing coins of different
// denominations (e.g. 1 dollar, 5 dollars, etc) and an integer amount
// representing a target amount of money.

// Return the fewest number of coins that you need to make up the
// exact target amount.
// If it is impossible to make up the amount, return -1.

// You may assume that you have an unlimited number of each coin.

// Example 1:
// Input: coins = [1,5,10], amount = 12
// Output: 3

// Explanation: 12 = 10 + 1 + 1.
// Note that we do not have to use every kind coin available.

// Take each amount
//.   Take each coin
//       Subtract that coin from the amount and see if I have already solved
//       the problem for that amoount, if yes:
//       => solution is 1 extra coin + solution for the sub amount

func coinChange(coins []int, amount int) int {
	// Create a memoization map to store the minimum number of coins needed for each amount
	// Key: amount, Value: minimum coins needed (or -1 if impossible)
	memo := map[int]int{
		0: 0, // Base case: 0 coins needed to make amount 0
	}

	// Build up solutions bottom-up: solve for amounts 1, 2, 3, ... up to target amount
	// This ensures when we calculate amount i, all smaller amounts are already solved
	for i := 1; i < amount+1; i++ {
		// Initialize minCoins to -1 (meaning "impossible to make this amount")
		// We'll update this if we find any valid way to make amount i
		minCoins := -1

		// Try using each coin denomination to make amount i
		for _, coin := range coins {
			// Skip this coin if it's larger than the current amount
			// (can't use a $5 coin to make $3)
			if coin > i {
				continue
			}

			// Calculate the subproblem: if we use this coin, what amount remains?
			// Example: if i=12 and coin=5, subproblem=7 (we need 7 more dollars)
			subproblem := i - coin

			// If subproblem is impossible (-1), skip this coin
			// Example: if i=8, coins=[5], coin=5, then subproblem=3 and that would be impossible
			subSolution := memo[subproblem]
			if subSolution == -1 {
				continue
			}

			// Calculate new solution: subproblem coins + 1 (for the current coin)
			// Example: if i=8, coins=[1, 4, 5], coin=4, subproblem=4
			//             subSolution[4]=1, newSolution=2
			newSolution := subSolution + 1

			// Update minCoins if this is our first valid solution OR if it's better than previous
			if minCoins == -1 || newSolution < minCoins {
				minCoins = newSolution
			}
		}

		// Store the result for amount i in our memo
		// This will be -1 if impossible, or the minimum coins needed
		memo[i] = minCoins
	}

	// Look up the answer for our target amount
	return memo[amount]
}
