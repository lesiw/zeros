package zeros

import "sync"

// Chan is a zero-valueable channel wrapper that auto-initializes on first use.
type Chan[T any] struct {
	once sync.Once
	ch   chan T
}

func (c *Chan[T]) init() {
	c.once.Do(func() {
		c.ch = make(chan T)
	})
}

// C returns the underlying channel.
func (c *Chan[T]) C() chan T {
	c.init()
	return c.ch
}

// Close closes the underlying channel.
func (c *Chan[T]) Close() {
	c.init()
	close(c.ch)
}
