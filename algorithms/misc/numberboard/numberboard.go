package numberboard

// You are given a linear number board with numbered positions.

// A player starts at a given starting position and rolls a die with D sides,
// where the die can produce any integer value from 1 to D.

// Some board positions contain teleporters.
// If the player lands on a teleporter position, they are immediately moved
// to another specified position. Teleporters may chain.

// Your task is to determine all final positions the player could end up on
// after a single die roll, considering every possible die value from 1 to D.
// (including any teleport destination)

// name:      "multiple teleporters across rolls",
// start:     5,
// diceSides: 4,
// teleporters: map[int]int{
// 	7: 12,
// 	9: 3,
// },
// expected: []int{6, 7, 12, 8, 9, 3},

// Example 2:
// start:     1,
// diceSides: 2,
// teleporters: map[int]int{
// 	2: 6,
// 	6: 3,
// 	3: 2,
// },
// expected: []int{2, 3, 6},

func ReachablePositions(start int, diceSides int, teleporters map[int]int) []int {
	reachable := map[int]struct{}{}

	for sideIndex := range diceSides {
		roll := sideIndex + 1

		landing := start + roll
		reachable[landing] = struct{}{}

		for {
			newLanding, ok := teleporters[landing]
			if !ok {
				break
			}
			if _, ok := reachable[newLanding]; ok {
				break
			}
			reachable[newLanding] = struct{}{}
			landing = newLanding
		}
	}

	output := make([]int, 0, len(reachable))
	for r := range reachable {
		output = append(output, r)
	}

	return output
}
