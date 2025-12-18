package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("shutting down gracefully")
}

func run() (err error) {
	ctx := context.Background()

	for {
		err = makeRequest(ctx)
		if err != nil {
			return err
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func makeRequest(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:3003/v2/course", http.NoBody)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(">>> resp: ", len(body))
	return nil
}
