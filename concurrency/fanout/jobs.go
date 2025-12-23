package fanout

import (
	"fmt"
	"sync"
	"time"
)

// worker processes messages from the input channel
func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Worker %d processing: %s\n", id, job)
		// Simulate some work
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Worker %d finished: %s\n", id, job)
	}
}

func FanOutJobs() {
	// Create a channel for jobs
	jobs := make(chan string, 10)

	// WaitGroup to track when all workers are done
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 3

	// Start workers (fan-out)
	for i := 1; i <= numWorkers; i++ {
		wg.Go(func() {
			worker(i, jobs)
		})
	}

	// Send jobs to the channel
	messages := []string{
		"task-1",
		"task-2",
		"task-3",
		"task-4",
		"task-5",
		"task-6",
	}

	for _, msg := range messages {
		jobs <- msg
	}

	// Close the jobs channel (no more jobs will be sent)
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed!")
}
