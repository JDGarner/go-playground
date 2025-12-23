package fanout

import (
	"fmt"
	"sync"
	"time"
)

// read from a single channel, fan out to many workers
// each worker reads one message
func FanOut(input <-chan string, workerCount int) {
	var wg sync.WaitGroup

	for i := range workerCount {
		wg.Go(func() {
			stringWorker(input, i)
		})
	}

	wg.Wait()
}

func stringWorker(input <-chan string, worker int) {
	for val := range input {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(">>> worker %d processed string: %v\n", worker, val)
	}
}
