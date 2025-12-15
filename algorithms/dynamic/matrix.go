package dynamic

type Pair struct {
	Row int
	Col int
}

// Count the number of unique paths from top left to bottom right
// only allowed to move down or right
func CountPaths(rows, cols int) int {
	memo := make(map[Pair]int)

	var dfs func(pair Pair) int

	dfs = func(pair Pair) int {
		if pair.Row == rows || pair.Col == cols { // out of bounds
			return 0
		}
		if count, ok := memo[pair]; ok {
			return count
		}
		if pair.Row == rows-1 && pair.Col == cols-1 { // end
			return 1
		}

		down := Pair{
			Row: pair.Row + 1,
			Col: pair.Col,
		}
		right := Pair{
			Row: pair.Row,
			Col: pair.Col + 1,
		}
		memo[pair] = dfs(down) + dfs(right)
		return memo[pair]
	}

	return dfs(Pair{0, 0})
}
