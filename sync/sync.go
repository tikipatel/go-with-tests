package sync

// Counter struct
type Counter struct {
	value int
}

// Value function
func (c *Counter) Value() int {
	return c.value
}

// Increment function
func (c *Counter) Increment() {
	c.value++
}
