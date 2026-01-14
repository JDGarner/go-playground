// A simple command-line tool for querying JSON log files.
//
// It loads log entries from one or more files in a directory, supports
// filtering by key-value pairs (such as service and level), and prints
// matching results in chronological order.
//
// Usage:
//
// go run main.go -path=/path/to/logs -query="service=auth level=error"
//
// Running `go run main.go` with no flags uses the default path ("logs")
// and an empty query (no filtering).
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type LogEntry struct {
	RawLine      string    `json:"-"` // Store original line for output
	RawTimestamp string    `json:"timestamp"`
	Timestamp    time.Time `json:"-"` // The parsed timestamp
	Service      string    `json:"service"`
	Level        string    `json:"level"`
	Message      string    `json:"message"`
}

type Query struct {
	Key   string
	Value string
}

func main() {
	path := flag.String("path", "logs", "the path to the directory containing the log files")
	query := flag.String("query", "", "the query to filter the logs by, e.g. service=auth")

	flag.Parse()

	err := run(*path, *query)
	if err != nil {
		fmt.Println("log parsing failed: ", err)
		os.Exit(1)
	}
}

func run(path, query string) error {
	queries, err := parseQueries(query)
	if err != nil {
		return fmt.Errorf("invalid query: %v", err)
	}

	entries, err := parseLogFiles(path)
	if err != nil {
		return fmt.Errorf("failed to parse log files: %v", err)
	}

	// Sort chronologically
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Before(entries[j].Timestamp)
	})

	for _, entry := range entries {
		if matchesQueries(queries, entry) {
			fmt.Println(entry.RawLine)
		}
	}

	return nil
}

func parseQueries(query string) (queries []Query, err error) {
	if query == "" {
		return queries, nil
	}

	queryStrings := strings.Fields(query)

	for _, queryString := range queryStrings {
		parts := strings.SplitN(queryString, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid query format: %s (expected key=value)", parts)
		}

		queries = append(queries, Query{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return queries, nil
}

func matchesQueries(queries []Query, entry LogEntry) bool {
	for _, query := range queries {
		switch query.Key {
		case "service":
			if entry.Service != query.Value {
				return false
			}
		case "level":
			if entry.Level != query.Value {
				return false
			}
		// More complex query logic could be added here, e.g. timestamp/message filtering
		default:
			return false
		}
	}

	return true
}

func parseLogFiles(path string) (entries []LogEntry, err error) {
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip any directories/non log files
		if d.IsDir() || filepath.Ext(path) != ".log" {
			return nil
		}

		fileEntries, err := parseFile(path)
		if err != nil {
			return err
		}
		entries = append(entries, fileEntries...)

		return nil
	})

	return entries, err
}

func parseFile(path string) (entries []LogEntry, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry LogEntry

		line := scanner.Text()
		err = json.Unmarshal([]byte(line), &entry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log line: %v", err)
		}

		timestamp, err := time.Parse(time.RFC3339, entry.RawTimestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp: %v", err)
		}
		entry.Timestamp = timestamp
		entry.RawLine = line

		entries = append(entries, entry)
	}
	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	return entries, nil
}
