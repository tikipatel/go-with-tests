package sync

import "sync"

// Counter struct
type Counter struct {
	mu    sync.Mutex
	value int
}

// Value function
func (c *Counter) Value() int {
	return c.value
}

// Increment function
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}
