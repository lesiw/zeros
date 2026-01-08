package zeros

import "sync"

// OnceValue is a zero-valueable wrapper that executes and caches the result
// of a function on first call.
//
// Unlike sync.OnceValue which returns a function, OnceValue is a struct type
// that can be used as a zero-value field in other structs.
//
// If f panics, Do will panic with the same value on every call.
type OnceValue[T any] struct {
	once  sync.Once
	valid bool
	p     any
	value T
}

// Do calls f and caches its return value on the first call.
// Subsequent calls return the cached value without calling f.
//
// If f panics, Do will panic with the same value on every call.
func (o *OnceValue[T]) Do(f func() T) T {
	o.once.Do(func() {
		defer func() {
			o.p = recover()
			if !o.valid {
				panic(o.p)
			}
		}()
		o.value = f()
		o.valid = true
	})
	if !o.valid {
		panic(o.p)
	}
	return o.value
}

// OnceValues is a zero-valueable wrapper that executes and caches the results
// of a function on first call.
//
// Unlike sync.OnceValues which returns a function, OnceValues is a struct type
// that can be used as a zero-value field in other structs.
//
// If f panics, Do will panic with the same value on every call.
type OnceValues[T1, T2 any] struct {
	once  sync.Once
	valid bool
	p     any
	v1    T1
	v2    T2
}

// Do calls f and caches its return values on the first call.
// Subsequent calls return the cached values without calling f.
//
// If f panics, Do will panic with the same value on every call.
func (o *OnceValues[T1, T2]) Do(f func() (T1, T2)) (T1, T2) {
	o.once.Do(func() {
		defer func() {
			o.p = recover()
			if !o.valid {
				panic(o.p)
			}
		}()
		o.v1, o.v2 = f()
		o.valid = true
	})
	if !o.valid {
		panic(o.p)
	}
	return o.v1, o.v2
}
