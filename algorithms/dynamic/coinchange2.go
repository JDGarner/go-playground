package dynamic

// You are given an integer array coins representing coins of different
// denominations and an integer amount representing a total amount of money.

// Return the number of combinations that make up that amount.
// If that amount of money cannot be made up by any combination of the coins,
// return 0.

// You may assume that you have an infinite number of each kind of coin.
// The answer is guaranteed to fit into a signed 32-bit integer.

// Example 1:
// Input: amount = 5, coins = [1,2,5]
// Output: 4

// Explanation: there are four ways to make up the amount:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1

//            5
//.  1        2        5
// 1 2 5    1 2 5    1 2 5

func change(amount int, coins []int) int {
	// Map of amount to the number of ways it can be made with given coins
	memo := map[int]int{
		0: 1, // Base case: the number of ways to make an amount of 0 is 1 (use 0 coins)
	}

	for _, coin := range coins {
		for i := 1; i < amount+1; i++ {
			subproblem := i - coin

			// if you can't reach i using that coin
			if subproblem < 0 {
				continue
			}

			memo[i] = memo[i] + memo[subproblem]
		}
	}

	return memo[amount]
}
