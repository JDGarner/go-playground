package fanin

import "sync"

func New(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range input {
				output <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
