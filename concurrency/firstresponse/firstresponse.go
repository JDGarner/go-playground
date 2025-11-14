package firstresponse

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Given any number of url strings, returns the response of the first responder
func Fetch(urls ...string) *http.Response {
	responses := make(chan *http.Response)

	for _, rawURL := range urls {
		go func() {
			parsedURL, err := url.Parse(rawURL)
			if err != nil {
				fmt.Println("failed to parse url: ", err)
				return
			}

			req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
			if err != nil {
				fmt.Println("failed to create request: ", err)
				return
			}

			start := time.Now()
			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("request failed: ", err)
				return
			}
			duration := time.Since(start)
			fmt.Printf("Request for %s took %v\n", rawURL, duration)

			responses <- resp
		}()
	}

	return <-responses
}

// Given any number of url strings, returns a channel to read the responses from
func FetchAll(urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	var wg sync.WaitGroup

	for _, rawURL := range urls {
		wg.Go(func() {
			parsedURL, err := url.Parse(rawURL)
			if err != nil {
				fmt.Println("failed to parse url: ", err)
				return
			}

			req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
			if err != nil {
				fmt.Println("failed to create request: ", err)
				return
			}

			start := time.Now()
			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("request failed: ", err)
				return
			}
			duration := time.Since(start)
			fmt.Printf("Request for %s took %v\n", rawURL, duration)

			responses <- resp
		})
	}

	go func() {
		wg.Wait()
		close(responses)
	}()

	return responses
}
