package helpers

type testCase struct {
	Input    []int
	Target   int
	Expected int
	Found    bool
}

var SearchTestCases = map[string]testCase{
	"search case 1 - first element": {
		Input:    []int{0, 1, 6, 7, 18},
		Target:   0,
		Expected: 0,
		Found:    true,
	},
	"search case 2 - last element": {
		Input:    []int{1, 1, 1, 2, 19},
		Target:   19,
		Expected: 4,
		Found:    true,
	},
	"search case 3 - middle element": {
		Input:    []int{100, 100, 101, 103, 104, 106, 107, 110, 111, 115, 115, 116, 116},
		Target:   106,
		Expected: 5,
		Found:    true,
	},
	"search case 4 - not found": {
		Input:    []int{1, 3, 5, 7, 9},
		Target:   17,
		Expected: -1,
		Found:    false,
	},
	"search case 5 - empty array": {
		Input:    []int{},
		Target:   5,
		Expected: -1,
		Found:    false,
	},
	"search case 6 - single element found": {
		Input:    []int{42},
		Target:   42,
		Expected: 0,
		Found:    true,
	},
	"search case 7 - single element not found": {
		Input:    []int{42},
		Target:   100,
		Expected: -1,
		Found:    false,
	},
	"search case 10 - negative elements": {
		Input:    []int{-10, -5, -3, -2, -1},
		Target:   -10,
		Expected: 0,
		Found:    true,
	},
	"search case 11 - large sorted array target near end": {
		Input: func() []int {
			arr := make([]int, 1000)
			for i := 0; i < 1000; i++ {
				arr[i] = i
			}
			return arr
		}(),
		Target:   995,
		Expected: 995,
		Found:    true,
	},
	"search case 12 - large sorted array not found": {
		Input: func() []int {
			arr := make([]int, 1000)
			for i := 0; i < 1000; i++ {
				arr[i] = i * 2
			}
			return arr
		}(),
		Target:   999,
		Expected: -1,
		Found:    false,
	},
}
