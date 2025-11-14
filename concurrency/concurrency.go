package concurrency

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/JDGarner/go-playground/concurrency/fanin"
	"github.com/JDGarner/go-playground/concurrency/firstresponse"
	"github.com/JDGarner/go-playground/concurrency/generator"
)

func FanInExample() {
	strings1 := generator.Strings("hello", "my", "name", "is", "ham")
	strings2 := generator.Strings("goodbye", "mr", "ham")
	strings3 := generator.Strings("1", "2", "3", "4", "5", "6", "7")

	combined := fanin.New(strings1, strings2, strings3)

	for val := range combined {
		fmt.Println(val)
	}
}

func TickerGeneratorExample() {
	secondTicker := generator.Ticker(time.Second)

	for range 4 {
		fmt.Println(<-secondTicker)
	}
}

func TickerWithDoneChannel() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()

	time.Sleep(4 * time.Second)
	done <- true
	close(done)
}

func FibGeneratorExample() {
	fib := generator.Fibonacci()
	for range 12 {
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

func CancellableExample() {
	// Cancel manually
	fmt.Println("With manual cancellation")
	ctx1, cancel := context.WithCancel(context.Background())
	ch1 := generator.Cancellable(ctx1)

	for range 8 {
		fmt.Println(<-ch1)
	}
	cancel()

	// Cancel with timeout
	fmt.Println("With timeout cancellation")
	ctx2, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ch2 := generator.Cancellable(ctx2)

	for v := range ch2 {
		fmt.Println(v)
	}
	cancel()

	// Cancel with deadline
	fmt.Println("With deadline cancellation")
	ctx3, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	ch3 := generator.Cancellable(ctx3)

	for v := range ch3 {
		fmt.Println(v)
	}
	cancel()
}

func FirstResponseExample() {
	resp := firstresponse.Fetch(
		fmt.Sprintf("https://swapi.dev/api/people/%d", rand.Intn(50)),
		"https://catfact.ninja/fact",
		"https://official-joke-api.appspot.com/random_joke",
		"https://api.breakingbadquotes.xyz/v1/quotes",
		"https://api.adviceslip.com/advice",
	)

	data, err := getResponseData(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("first returned data:")
	fmt.Println(data)
}

func AllResponsesExample() {
	responses := firstresponse.FetchAll(
		fmt.Sprintf("https://swapi.dev/api/people/%d", rand.Intn(50)),
		"https://catfact.ninja/fact",
		"https://official-joke-api.appspot.com/random_joke",
		"https://api.breakingbadquotes.xyz/v1/quotes",
		"https://api.adviceslip.com/advice",
	)

	for resp := range responses {
		data, err := getResponseData(resp)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	}
}

func getResponseData(resp *http.Response) (string, error) {
	if resp == nil {
		return "", errors.New("no response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: ", err)
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal body into json: ", err)
	}

	return buf.String(), nil
}
