package searchwords

import "fmt"

func SearchWordsExample() {
	wd := Constructor()
	wd.AddWord("day")
	wd.AddWord("bay")
	wd.AddWord("may")
	res := wd.Search("say") // return false
	fmt.Println(">>> say ", res)
	res = wd.Search("day") // return true
	fmt.Println(">>> day ", res)
	res = wd.Search(".ay") // return true
	fmt.Println(">>> .ay ", res)
	res = wd.Search("b..") // return true
	fmt.Println(">>> b.. ", res)
	res = wd.Search(".a.") // return true
	fmt.Println(">>> .a. ", res)
	res = wd.Search(".b.") // return false
	fmt.Println(">>> .b. ", res)
	res = wd.Search("...") // return true
	fmt.Println(">>> ... ", res)
}
