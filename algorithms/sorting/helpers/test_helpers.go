package helpers

type testCase struct {
	Input    []int
	Expected []int
}

var SortingTestCases = map[string]testCase{
	"sorting case 1": {
		Input:    []int{6, 18, 3, 1, 7, 29, 2, 0, 4},
		Expected: []int{0, 1, 2, 3, 4, 6, 7, 18, 29},
	},
	"sorting case 2": {
		Input:    []int{1, 1, 19, 2, 1},
		Expected: []int{1, 1, 1, 2, 19},
	},
	"sorting case 3": {
		Input:    []int{42},
		Expected: []int{42},
	},
	"sorting case 4": {
		Input:    []int{42, 7},
		Expected: []int{7, 42},
	},
	"sorting case 5 - negative numbers": {
		Input:    []int{-3, -1, -7, 2, 0},
		Expected: []int{-7, -3, -1, 0, 2},
	},
	"sorting case 6 - already sorted": {
		Input:    []int{1, 2, 3, 4, 5},
		Expected: []int{1, 2, 3, 4, 5},
	},
	"sorting case 7 - reverse sorted": {
		Input:    []int{9, 7, 5, 3, 1},
		Expected: []int{1, 3, 5, 7, 9},
	},
	"sorting case 8 - duplicates and negatives": {
		Input:    []int{-1, 5, 3, -1, 5},
		Expected: []int{-1, -1, 3, 5, 5},
	},
	"sorting case 9 - all equal elements": {
		Input:    []int{8, 8, 8, 8},
		Expected: []int{8, 8, 8, 8},
	},
	"sorting case 10 - elements with many duplicates": {
		Input:    []int{6, 4, 1, 2, 3, 9, 10, 5, 6, 4, 4, 2, 3, 1, 5, 8, 3, 2, 9, 10, 1},
		Expected: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 6, 6, 8, 9, 9, 10, 10},
	},
	"sorting case 11 - range not starting from 1": {
		Input:    []int{100, 101, 116, 107, 111, 115, 115, 110, 106, 103, 100, 116, 104},
		Expected: []int{100, 100, 101, 103, 104, 106, 107, 110, 111, 115, 115, 116, 116},
	},
}
