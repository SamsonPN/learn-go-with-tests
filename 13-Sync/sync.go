package main

import "sync"

type Counter struct {
	// mutex = mutual exclusion lock
	// this should NEVER be embedded into the struct
	// i.e. it should never be called anonymously
	// reason being, the mutex becomes public!!!
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// lock ensures that each goroutine does not
	// increment at the same time
	// the goroutine that has the lock will block other goroutines
	// from running and they have to wait for it to be unlocked
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
