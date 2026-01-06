package intervals

import "fmt"

func MergeIntervalsExample() {
	// intervals := [][]int{
	// 	{1, 3}, {1, 5}, {6, 7},
	// }
	// intervals := [][]int{
	// 	{8, 10}, {15, 18},
	// }
	// intervals := [][]int{
	// 	{0, 2}, {1, 4}, {3, 5},
	// }

	// intervals := [][]int{
	// 	{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	// }
	// res := insert(intervals, []int{4, 8})

	// [[1,5]]
	// [6,8]

	// intervals := [][]int{{0, 5}}
	// toBeRemoved := []int{2, 3}
	// expected:    [][]int{{0, 2}, {3, 5}},

	// intervals := [][]int{{5, 10}}
	// toBeRemoved := []int{5, 7}
	// expected:    [][]int{{7, 10}},

	// intervals := [][]int{{5, 10}}
	// toBeRemoved := []int{5, 10}
	// // expected:    [][]int{},

	// res := remove(intervals, toBeRemoved)

	// fmt.Println(">>> res: ", res)

	// res := AddBoldTag("abcxyz123", []string{"abc", "123"})
	// // expect: "<b>abc</b>xyz<b>123</b>",
	// res := AddBoldTag("aaabbb", []string{"aa", "b"})
	// expect: "<b>aaabbb</b>",
	res := AddBoldTag("你好ab", []string{"好", "ab"})
	// expect: "你<b>好</b><b>ab</b>",

	fmt.Println(">>> res: ", res)

}
