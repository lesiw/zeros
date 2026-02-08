package zeros

import "iter"

// Map is a zero-valueable map wrapper that auto-initializes on first use.
type Map[K comparable, V any] struct{ once OnceValue[map[K]V] }

func (m *Map[K, V]) init() map[K]V { return make(map[K]V) }

// Map returns the underlying map.
func (m *Map[K, V]) Map() map[K]V { return m.once.Do(m.init) }

// Set sets a key-value pair.
func (m *Map[K, V]) Set(key K, value V) { m.once.Do(m.init)[key] = value }

// Get retrieves a value by key, returning the zero value if not found.
func (m *Map[K, V]) Get(key K) V { return m.once.Do(m.init)[key] }

// CheckGet retrieves a value by key with a presence indicator.
func (m *Map[K, V]) CheckGet(key K) (V, bool) {
	v, ok := m.once.Do(m.init)[key]
	return v, ok
}

// Delete removes a key.
func (m *Map[K, V]) Delete(key K) { delete(m.once.Do(m.init), key) }

// Len returns the number of elements.
func (m *Map[K, V]) Len() int { return len(m.once.Do(m.init)) }

// Keys returns an iterator over keys in the map.
func (m *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m.once.Do(m.init) {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns an iterator over values in the map.
func (m *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m.once.Do(m.init) {
			if !yield(v) {
				return
			}
		}
	}
}

// All returns an iterator over key-value pairs in the map.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m.once.Do(m.init) {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() { clear(m.once.Do(m.init)) }
