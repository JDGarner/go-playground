package hashmap

import "fmt"

func Example() {
	h := New[string]()

	h.Insert("carbon", "6")
	h.Insert("iron", "26")
	h.Insert("bromine", "35")
	h.Insert("xenon", "54")

	fmt.Println(h)

	h.Insert("moscovium", "115")
	h.Insert("cobalt", "27")
	h.Insert("wolfram", "74")
	h.Insert("curium", "wrong-value")
	h.Insert("curium", "96")

	fmt.Println(h)

	carbon, ok := h.Get("carbon")
	if ok {
		fmt.Println("carbon: ", carbon)
	} else {
		fmt.Println("carbon not found!")
	}

	moscovium, ok := h.Get("moscovium")
	if ok {
		fmt.Println("moscovium: ", moscovium)
	} else {
		fmt.Println("moscovium not found!")
	}

	plutonium, ok := h.Get("plutonium")
	if ok {
		fmt.Println("plutonium: ", plutonium)
	} else {
		fmt.Println("plutonium not found!")
	}
}
