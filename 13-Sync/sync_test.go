package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount) // sets the number of goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // basically wg.Counter--
			}()
		}

		// waits until all of the goroutines
		// are finished before proceeding
		wg.Wait()

		// Mutex must not be passed by value where it is copied
		// therefore, we should use pointers
		assertCounter(t, counter, wantedCount)
	})
}

// since we just want to use the pointer
// due to the Mutex
// it is better to make our constructor
// to return the pointer
// and shows the reader that it would be better
// to not initialize the type themselves!
func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
