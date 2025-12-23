package fanout

import (
	"sync"
)

// read from a single channel, fan out to many channels
// each returned channel reads all the messages from input
func FanOutBroadcast(input <-chan string, count int) []chan string {
	outputs := make([]chan string, count)
	// Initialize all output channels
	for i := range outputs {
		outputs[i] = make(chan string)
	}

	var wg sync.WaitGroup

	wg.Go(func() {
		for val := range input {
			for _, output := range outputs {
				wg.Go(func() {
					output <- val
				})
			}
		}
	})

	go func() {
		wg.Wait()
		for _, output := range outputs {
			close(output)
		}
	}()

	return outputs
}
