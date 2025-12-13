package errgroup

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Example() {
	ctx := context.Background()
	g, ctx := WithContext(ctx)

	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://invalid.host", // will cause error
	}

	client := http.Client{Timeout: 5 * time.Second}

	results := make([]string, len(urls))

	for i, url := range urls {
		i, url := i, url
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return err
			}

			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			results[i] = resp.Status
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("results:", results)
}
