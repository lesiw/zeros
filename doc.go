// Package zeros provides zero-valueable wrappers for channels, maps,
// and sync.OnceValues.
//
// Chan and Map auto-initialize their underlying types on first use,
// allowing them to be used without explicit initialization.
// Initialization is thread-safe, but the types themselves are not safe
// for concurrent access without external synchronization
// (like Go's built-in map type).
//
// OnceValue and OnceValues provide zero-valueable alternatives to
// sync.OnceValue and sync.OnceValues, allowing them to be used as
// struct fields without initialization.
//
// Example usage:
//
//	var ch zeros.Chan[int]
//	ch.C() <- 42  // auto-initializes the channel
//	v := <-ch.C()
//
//	var m zeros.Map[string, int]
//	m.Set("answer", 42)  // auto-initializes the map
//	v, ok := m.Get("answer")
//
//	for k, v := range m.All() {
//		fmt.Printf("%s: %d\n", k, v)
//	}
//
//	type Config struct {
//		value zeros.OnceValue[int]
//	}
//	result := config.value.Do(func() int { return 42 })
package zeros
