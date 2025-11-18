package fanin

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	numLogFiles = 3
)

func Logger() {
	for i := range numLogFiles {
		go writer(fmt.Sprintf("./temp-logs-%d.log", i))
	}

	time.Sleep(1 * time.Second) // Wait for files to be created

	var readers []<-chan string

	for i := range numLogFiles {
		r, err := reader(fmt.Sprintf("./temp-logs-%d.log", i))
		if err != nil {
			panic(err)
		}

		readers = append(readers, r)
	}

	merged := New(readers...)

	go func() {
		for value := range merged {
			fmt.Println(value)
		}
	}()

	select {}
}

func writer(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer f.Close()

	logger := log.New(f, "INFO: ", log.Ldate|log.Ltime)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		logger.Printf("Periodic log entry for logger: %s", path)
	}
}

func reader(path string) (<-chan string, error) {
	lines := make(chan string)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	go func() {
		defer file.Close()
		defer close(lines)

		r := bufio.NewReader(file)

		for {
			line, err := r.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					time.Sleep(200 * time.Millisecond)
					continue
				}
				return
			}
			lines <- strings.TrimRight(line, "\n")
		}
	}()

	return lines, nil
}
