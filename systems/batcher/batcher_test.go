package batcher_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/JDGarner/go-playground/systems/batcher"
	"github.com/stretchr/testify/assert"
)

type API struct {
	processed atomic.Int32
}

type RequestParams struct {
	ID int
}

type Result struct {
	ID int
}

func (a *API) Process(params []RequestParams) batcher.ProcessResult[Result] {
	a.processed.Add(int32(len(params)))

	return batcher.ProcessResult[Result]{
		Result: Result{
			ID: 1,
		},
		Err: nil,
	}
}

func TestBatcher(t *testing.T) {
	t.Run("processes batch after waitTime has elapsed", func(t *testing.T) {
		a := &API{}

		b := batcher.New(10, 500*time.Millisecond, a)

		for i := range 10 {
			b.Add(RequestParams{ID: i})
		}

		assert.Zero(t, a.processed.Load())
		assert.Eventually(t, func() bool {
			return a.processed.Load() == 10
		}, 600*time.Millisecond, 50*time.Millisecond)
	})

	t.Run("only proccesses batchSize number of jobs at a time", func(t *testing.T) {
		a := &API{}

		b := batcher.New(10, 500*time.Millisecond, a)

		for i := range 15 {
			b.Add(RequestParams{ID: i})
		}

		assert.Zero(t, a.processed.Load())
		assert.Eventually(t, func() bool {
			return a.processed.Load() == 10
		}, 600*time.Millisecond, 50*time.Millisecond)
	})

	t.Run("processes all remaining jobs in queue before shutting down", func(t *testing.T) {
		a := &API{}

		b := batcher.New(10, 50*time.Millisecond, a)

		for i := range 30 {
			b.Add(RequestParams{ID: i})
		}

		assert.Zero(t, a.processed.Load())

		b.Close()

		assert.EqualValues(t, a.processed.Load(), 30)
	})

	t.Run("processes all remaining jobs in queue before shutting down", func(t *testing.T) {
		a := &API{}

		b := batcher.New(10, 50*time.Millisecond, a)

		for i := range 30 {
			b.Add(RequestParams{ID: i})
		}

		assert.Zero(t, a.processed.Load())

		b.Close()

		assert.EqualValues(t, a.processed.Load(), 30)
	})
}
