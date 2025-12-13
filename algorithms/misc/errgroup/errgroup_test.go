package errgroup

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestErrGroup_NoErrors(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	var counter atomic.Int32

	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			counter.Add(1)
			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if counter.Load() != 10 {
		t.Errorf("expected counter to be 10, got %d", counter.Load())
	}

	// Context should be cancelled after Wait()
	select {
	case <-ctx.Done():
		// expected
	default:
		t.Error("context should be cancelled after Wait()")
	}
}

func TestErrGroup_SingleError(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	expectedErr := errors.New("test error")

	eg.Go(func() error {
		return expectedErr
	})

	eg.Go(func() error {
		return nil
	})

	err := eg.Wait()
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	// Context should be cancelled
	select {
	case <-ctx.Done():
		// expected
	default:
		t.Error("context should be cancelled when error occurs")
	}
}

func TestErrGroup_MultipleErrors(t *testing.T) {
	eg, _ := WithContext(context.Background())

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	eg.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return err1
	})

	eg.Go(func() error {
		time.Sleep(20 * time.Millisecond)
		return err2
	})

	err := eg.Wait()
	// Should get the first error (err1 since it completes first)
	if err != err1 {
		t.Errorf("expected error %v, got %v", err1, err)
	}
}

func TestErrGroup_ContextCancellation(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	var cancelledCount atomic.Int32

	for i := 0; i < 5; i++ {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				cancelledCount.Add(1)
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		})
	}

	// One goroutine that errors quickly
	eg.Go(func() error {
		return errors.New("quick error")
	})

	err := eg.Wait()
	if err == nil {
		t.Error("expected an error")
	}

	// Give goroutines time to detect cancellation
	time.Sleep(50 * time.Millisecond)

	// At least some goroutines should have been cancelled
	if cancelledCount.Load() == 0 {
		t.Error("expected some goroutines to be cancelled")
	}
}

func TestErrGroup_ParentContextCancelled(t *testing.T) {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	eg, ctx := WithContext(parentCtx)

	started := make(chan struct{})
	var gotContextError atomic.Bool

	eg.Go(func() error {
		close(started)
		<-ctx.Done()
		gotContextError.Store(true)
		return ctx.Err()
	})

	// Wait for goroutine to start
	<-started

	// Cancel parent context
	parentCancel()

	err := eg.Wait()
	if err == nil {
		t.Error("expected context cancellation error")
	}

	if !gotContextError.Load() {
		t.Error("goroutine should have received context cancellation")
	}

	if !errors.Is(err, context.Canceled) {
		t.Error("error should be context canceled")
	}
}

func TestErrGroup_ErrorPropagation(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	testErr := errors.New("specific error")

	eg.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return testErr
	})

	// This should be cancelled before completing
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
			t.Error("goroutine should have been cancelled")
			return nil
		}
	})

	err := eg.Wait()

	// Should get either the specific error or a context error
	if err != testErr && !errors.Is(err, context.Canceled) {
		t.Errorf("expected %v or context.Canceled, got %v", testErr, err)
	}
}

func TestErrGroup_ConcurrentGoCalls(t *testing.T) {
	eg, _ := WithContext(context.Background())

	var counter atomic.Int32
	numGoroutines := 100

	// Spawn goroutines concurrently
	var startWg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		startWg.Add(1)
		go func() {
			startWg.Done()
			eg.Go(func() error {
				counter.Add(1)
				return nil
			})
		}()
	}

	startWg.Wait()
	err := eg.Wait()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if counter.Load() != int32(numGoroutines) {
		t.Errorf("expected counter to be %d, got %d", numGoroutines, counter.Load())
	}
}

func TestErrGroup_NoGoRoutines(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	err := eg.Wait()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Context should still be cancelled
	select {
	case <-ctx.Done():
		// expected
	default:
		t.Error("context should be cancelled after Wait()")
	}
}

func TestErrGroup_GoAfterError(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	expectedErr := errors.New("first error")

	eg.Go(func() error {
		return expectedErr
	})

	// This should still be registered even though an error occurred
	eg.Go(func() error {
		<-ctx.Done()
		return nil
	})

	err := eg.Wait()
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func TestErrGroup_LongRunningTasks(t *testing.T) {
	eg, ctx := WithContext(context.Background())

	var completed atomic.Int32

	// Long running task
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(500 * time.Millisecond):
			completed.Add(1)
			return nil
		}
	})

	// Quick error
	eg.Go(func() error {
		time.Sleep(10 * time.Millisecond)
		return errors.New("quick error")
	})

	start := time.Now()
	err := eg.Wait()
	duration := time.Since(start)

	if err == nil {
		t.Error("expected an error")
	}

	// Should complete quickly due to cancellation, not wait 500ms
	if duration > 200*time.Millisecond {
		t.Errorf("Wait() took too long: %v", duration)
	}

	// Long running task should have been cancelled
	if completed.Load() != 0 {
		t.Error("long running task should have been cancelled")
	}
}
