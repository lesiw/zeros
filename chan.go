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

// Chan returns the underlying channel.
func (c *Chan[T]) Chan() chan T {
	c.init()
	return c.ch
}

// Send sends a value on the channel, blocking until the value is sent.
func (c *Chan[T]) Send(v T) {
	c.init()
	c.ch <- v
}

// Recv receives a value from the channel, blocking until a value is available.
// Returns the zero value if the channel is closed.
func (c *Chan[T]) Recv() T {
	c.init()
	v := <-c.ch
	return v
}

// CheckRecv receives a value from the channel with a status indicator.
// The boolean return value indicates whether the channel is open.
func (c *Chan[T]) CheckRecv() (T, bool) {
	c.init()
	v, ok := <-c.ch
	return v, ok
}

// TrySend attempts to send a value on the channel without blocking.
// It reports whether the value was sent.
func (c *Chan[T]) TrySend(v T) bool {
	c.init()
	select {
	case c.ch <- v:
		return true
	default:
		return false
	}
}

// TryRecv attempts to receive a value from the channel without blocking.
// If a value is available, it returns the value and true.
// If no value is available, it returns the zero value and false.
// If the channel is closed, it returns the zero value and false.
func (c *Chan[T]) TryRecv() (T, bool) {
	c.init()
	select {
	case v, ok := <-c.ch:
		return v, ok
	default:
		var zero T
		return zero, false
	}
}

// Close closes the underlying channel.
func (c *Chan[T]) Close() {
	c.init()
	close(c.ch)
}
