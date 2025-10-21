package generator

import (
	"context"
	"time"
)

func Integer(num int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			ch <- i
		}
	}()

	return ch
}

func Double(inputCh <-chan int) <-chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)
		for input := range inputCh {
			outputCh <- input * 2
		}
	}()

	return outputCh
}

func Fibonacci() <-chan int {
	fibCh := make(chan int)

	go func() {
		a, b := 0, 1

		for {
			fibCh <- a
			a, b = b, a+b
		}
	}()

	return fibCh
}

func Strings(strings ...string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for _, s := range strings {
			ch <- s
		}
	}()

	return ch
}

func Ticker(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)

	go func() {
		t := time.NewTicker(d)
		defer t.Stop()
		defer close(ch)

		for tick := range t.C {
			ch <- tick
		}
	}()

	return ch
}

func Cancellable(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; ; i++ {
			select {
			case ch <- i:
				time.Sleep(time.Millisecond * 200)
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}
