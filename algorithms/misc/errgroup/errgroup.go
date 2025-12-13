package errgroup

import (
	"context"
	"sync"
)

type ErrGroup struct {
	ctx    context.Context
	cancel context.CancelFunc
	errCh  chan error
	wg     sync.WaitGroup
}

// cancels all goroutines when any returns an error
func WithContext(c context.Context) (*ErrGroup, context.Context) {
	ctx, cancel := context.WithCancel(c)

	return &ErrGroup{
		ctx:    ctx,
		cancel: cancel,
		errCh:  make(chan error, 1),
	}, ctx
}

// Registers the func as a concurrent task
func (e *ErrGroup) Go(f func() error) {
	e.wg.Go(func() {
		if err := f(); err != nil {
			select {
			case e.errCh <- err:
				e.cancel()
			default:
				// If errCh is already full we go here immediately - first error already recorded
			}
		}
	})
}

// Blocks until all go routines are done
// Returns the first error from go routine (or from ctx cancellation) if there is any
func (e *ErrGroup) Wait() (err error) {
	defer e.cancel()

	done := make(chan struct{})

	// If wg.Wait returns, close the done channel
	go func() {
		e.wg.Wait()
		close(done)
	}()

	// Either:
	// - read error from errCh then wait for done and return error
	// - OR wait for done then return nil
	// - OR wait for ctx cancellation, wait for done, then return ctx error
	select {
	case err := <-e.errCh:
		<-done
		return err
	case <-done:
		return nil
	case <-e.ctx.Done():
		<-done
		return e.ctx.Err()
	}
}
