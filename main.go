package main

import (
	"fmt"

	fanin "github.com/JDGarner/go-playground/concurrency/fan_in"
	"github.com/JDGarner/go-playground/concurrency/generator"
)

func main() {
	// FanInExample()
	// FibGeneratorExample()
	// DoubleGeneratorExample()
}

func FanInExample() {
	strings1 := generator.Strings("hello", "my", "name", "is", "ham")
	strings2 := generator.Strings("goodbye", "mr", "ham")

	combined := fanin.New(strings1, strings2)

	for val := range combined {
		fmt.Println(val)
	}
}

func FibGeneratorExample() {
	fib := generator.Fibonacci()
	for i := 0; i < 12; i++ {
		fmt.Println(<-fib)
	}
}

func DoubleGeneratorExample() {
	numbers := generator.Integer(10)
	doubled := generator.Double(numbers)

	for num := range doubled {
		fmt.Println(num)
	}
}
