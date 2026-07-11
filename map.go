package zeros

import "iter"

// Map is a zero-valueable map wrapper that auto-initializes on first use.
type Map[K comparable, V any] struct{ once OnceValue[map[K]V] }

// Map returns the underlying map.
func (m *Map[K, V]) Map() map[K]V {
	return m.once.Do(func() map[K]V { return make(map[K]V) })
}

// Set sets a key-value pair.
func (m *Map[K, V]) Set(key K, value V) { m.Map()[key] = value }

// Get retrieves a value by key, returning the zero value if not found.
func (m *Map[K, V]) Get(key K) V { return m.Map()[key] }

// CheckGet retrieves a value by key with a presence indicator.
func (m *Map[K, V]) CheckGet(key K) (V, bool) {
	v, ok := m.Map()[key]
	return v, ok
}

// Delete removes a key.
func (m *Map[K, V]) Delete(key K) { delete(m.Map(), key) }

// Len returns the number of elements.
func (m *Map[K, V]) Len() int { return len(m.Map()) }

// Keys returns an iterator over keys in the map.
func (m *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m.Map() {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns an iterator over values in the map.
func (m *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m.Map() {
			if !yield(v) {
				return
			}
		}
	}
}

// All returns an iterator over key-value pairs in the map.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m.Map() {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() { clear(m.Map()) }
