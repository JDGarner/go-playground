package voteballot

// Design a system to count votes from ballots where:
// Each ballot can rank up to 3 candidates.

// Point system:
// 1st place → 3 points
// 2nd place → 2 points
// 3rd place → 1 point

// Return the winning candidate by total points.

type Ballot struct {
	id     string
	first  Vote
	second Vote
	third  Vote
}

const (
	ThirdPlaceVote  = iota + 1 // 1
	SecondPlaceVote            // 2
	FirstPlaceVote             // 3
)

type Vote struct {
	candidate Candidate
}

type Candidate struct {
	name string
}

type Winner struct {
	candidate Candidate
	votes     int
}

type TieBreaker interface {
	BreakTie(ballots []Ballot, tied []Winner) Winner
}

// Return the winning candidate by total points.
func CountVotes(ballots []Ballot, tb TieBreaker) (winner Winner) {
	results := make(map[Candidate]int)
	counted := make(map[string]struct{})

	for _, ballot := range ballots {
		if _, ok := counted[ballot.id]; ok {
			continue
		}
		counted[ballot.id] = struct{}{}

		countBallot(results, ballot.first.candidate, FirstPlaceVote)
		countBallot(results, ballot.second.candidate, SecondPlaceVote)
		countBallot(results, ballot.third.candidate, ThirdPlaceVote)
	}

	max := 0

	var winners []Winner

	for current, voteCount := range results {
		if voteCount > max {
			max = voteCount
			winners = []Winner{
				{
					candidate: current,
					votes:     voteCount,
				},
			}
		} else if voteCount == max {
			winners = append(winners, Winner{
				candidate: current,
				votes:     voteCount,
			})
		}
	}

	if len(winners) == 1 {
		return winners[0]
	}

	return tb.BreakTie(ballots, winners)
}

func countBallot(results map[Candidate]int, candidate Candidate, score int) {
	if _, ok := results[candidate]; !ok {
		results[candidate] = 0
	}
	results[candidate] += score
}

// Part 2
// Implement a tie-breaking strategy:

// If multiple candidates have the same maximum points,
// the winner is the one who reached the max points first
// in the tallying order (time-series based).

func FirstToReachMax(ballots []Ballot, tiedWinners []Winner) Winner {
	target := tiedWinners[0].votes
	results := make(map[Candidate]int)

	for _, ballot := range ballots {
		countBallot(results, ballot.first.candidate, FirstPlaceVote)
		if w := checkTieBreaker(results, ballot.first.candidate, target); w != nil {
			return *w
		}
		countBallot(results, ballot.second.candidate, SecondPlaceVote)
		if w := checkTieBreaker(results, ballot.second.candidate, target); w != nil {
			return *w
		}
		countBallot(results, ballot.third.candidate, ThirdPlaceVote)
		if w := checkTieBreaker(results, ballot.third.candidate, target); w != nil {
			return *w
		}
	}

	return Winner{} // won't reach here
}

func checkTieBreaker(results map[Candidate]int, candidate Candidate, target int) *Winner {
	if results[candidate] >= target {
		return &Winner{
			candidate: candidate,
			votes:     target,
		}
	}

	return nil
}
