package array

func groupAnagrams(strs []string) [][]string {
	alphabet := map[rune]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
		'i': 8,
		'j': 9,
		'k': 10,
		'l': 11,
		'm': 12,
		'n': 13,
		'o': 14,
		'p': 15,
		'q': 16,
		'r': 17,
		's': 18,
		't': 19,
		'u': 20,
		'v': 21,
		'w': 22,
		'x': 23,
		'y': 24,
		'z': 25,
	}

	groupedWords := make(map[[26]int][]string)

	for _, s := range strs {
		// create freq array such that freq[alphabetIndex] = freq count of that letter
		var freq [26]int

		for _, b := range s {
			alphabetIndex := alphabet[b]
			freq[alphabetIndex]++
		}

		if _, ok := groupedWords[freq]; !ok {
			groupedWords[freq] = []string{}
		}
		groupedWords[freq] = append(groupedWords[freq], s)
	}

	res := [][]string{}

	for _, group := range groupedWords {
		res = append(res, group)
	}

	return res
}

func groupAnagrams2(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
