package zeros

import (
	"iter"
	"sync"
)

// Map is a zero-valueable map wrapper that auto-initializes on first use.
type Map[K comparable, V any] struct {
	once sync.Once
	m    map[K]V
}

func (m *Map[K, V]) init() {
	m.once.Do(func() {
		m.m = make(map[K]V)
	})
}

// Map returns the underlying map.
func (m *Map[K, V]) Map() map[K]V {
	m.init()
	return m.m
}

// Set sets a key-value pair.
func (m *Map[K, V]) Set(key K, value V) {
	m.init()
	m.m[key] = value
}

// Get retrieves a value by key, returning the zero value if not found.
func (m *Map[K, V]) Get(key K) V {
	if m.m == nil {
		var zero V
		return zero
	}
	return m.m[key]
}

// CheckGet retrieves a value by key with a presence indicator.
func (m *Map[K, V]) CheckGet(key K) (V, bool) {
	if m.m == nil {
		var zero V
		return zero, false
	}
	v, ok := m.m[key]
	return v, ok
}

// Delete removes a key.
func (m *Map[K, V]) Delete(key K) {
	if m.m == nil {
		return
	}
	delete(m.m, key)
}

// Len returns the number of elements.
func (m *Map[K, V]) Len() int {
	if m.m == nil {
		return 0
	}
	return len(m.m)
}

// Keys returns an iterator over keys in the map.
func (m *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		if m.m == nil {
			return
		}
		for k := range m.m {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns an iterator over values in the map.
func (m *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		if m.m == nil {
			return
		}
		for _, v := range m.m {
			if !yield(v) {
				return
			}
		}
	}
}

// All returns an iterator over key-value pairs in the map.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m.m == nil {
			return
		}
		for k, v := range m.m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.init()
	clear(m.m)
}
