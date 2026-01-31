package batcher

import (
	"sync"
	"time"
)

type BatchProcessor[T any, R any] interface {
	Process(jobs []T) ProcessResult[R]
}

type ProcessResult[R any] struct {
	Result R
	Err    error
}

type Batcher[T any, R any] struct {
	batchSize int
	waitTime  time.Duration
	processor BatchProcessor[T, R]
	queue     []T
	mu        sync.Mutex
	wg        sync.WaitGroup
	closed    bool
	results   chan (ProcessResult[R])
}

func New[T any, R any](batchSize int, waitTime time.Duration, processor BatchProcessor[T, R]) *Batcher[T, R] {
	b := &Batcher[T, R]{
		batchSize: batchSize,
		waitTime:  waitTime,
		processor: processor,
		results:   make(chan ProcessResult[R], 100), // add a buffer so results don't have to be read immediately
	}

	b.wg.Go(b.run)

	return b
}

func (b *Batcher[T, R]) run() {
	ticker := time.NewTicker(b.waitTime)
	defer ticker.Stop()

	for range ticker.C {
		b.process()
		if b.shouldEndRun() {
			return
		}
	}
}

func (b *Batcher[T, R]) shouldEndRun() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.closed && len(b.queue) == 0
}

func (b *Batcher[T, R]) process() {
	b.mu.Lock()
	defer b.mu.Unlock()

	// take batchSize number from the queue
	num := b.batchSize
	if len(b.queue) < num {
		num = len(b.queue)
	}

	batch := make([]T, num)
	copy(batch, b.queue[:num])
	b.queue = b.queue[num:]

	b.wg.Go(func() {
		b.results <- b.processor.Process(batch)
	})
}

func (b *Batcher[T, R]) Add(job T) bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return false
	}

	b.queue = append(b.queue, job)
	return true
}

func (b *Batcher[T, R]) Results() <-chan ProcessResult[R] {
	return b.results
}

func (b *Batcher[T, R]) Close() {
	b.mu.Lock()
	b.closed = true
	b.mu.Unlock()

	// wait until ALL jobs in the queue have been processed
	b.wg.Wait()
}
