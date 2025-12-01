# zeros

[![Go Reference](https://pkg.go.dev/badge/lesiw.io/zeros.svg)](https://pkg.go.dev/lesiw.io/zeros)

Zero-valueable wrappers for Go channels and maps.

## Overview

Package `zeros` provides `Chan[T]` and `Map[K,V]` types that auto-initialize on first use, eliminating the need for explicit `make()` calls. This makes them usable at their zero value, following the same principle as types like `bytes.Buffer` and `sync.Mutex`.

## Installation

```bash
go get lesiw.io/zeros
```

## Features

- **Zero-value usability**: Use `var ch zeros.Chan[int]` or `var m zeros.Map[string,int]` without initialization
- **Thread-safe initialization**: Initialization is thread-safe via `sync.Once`
- **Minimal API**: Simple wrappers that mirror built-in types
- **Modern Go**: Uses Go 1.23+ features including range-over-func iterators

## Usage

### Chan

[▶️ Run this example on the Go Playground](https://go.dev/play/p/PUEddCMEVAR)

```go
package main

import (
    "fmt"

    "lesiw.io/zeros"
)

func main() {
    var ch zeros.Chan[int]

    go func() {
        ch.Send(42) // auto-initializes the channel
    }()

    value := ch.Recv()
    fmt.Println(value)
}
```

Available methods:
- `Chan() chan T` - Returns the underlying channel
- `Send(v T)` - Sends a value on the channel (blocks)
- `Recv() T` - Receives a value from the channel (blocks)
- `CheckRecv() (T, bool)` - Receives with channel status indicator
- `TrySend(v T) bool` - Attempts to send without blocking
- `TryRecv() (T, bool)` - Attempts to receive without blocking
- `Close()` - Closes the underlying channel

### Map

[▶️ Run this example on the Go Playground](https://go.dev/play/p/JahXu_ZR9DR)

```go
package main

import (
    "fmt"

    "lesiw.io/zeros"
)

func main() {
    var m zeros.Map[string, int]

    m.Set("answer", 42) // auto-initializes the map

    fmt.Println(m.Get("answer"))

    if _, ok := m.CheckGet("missing"); !ok {
        fmt.Println("key not found")
    }

    for k, v := range m.All() {
        fmt.Println(k, v)
    }
}
```

Available methods:
- `Map() map[K]V` - Returns the underlying map
- `Set(key K, value V)` - Sets a key-value pair
- `Get(key K) V` - Returns value or zero value if missing
- `CheckGet(key K) (V, bool)` - Returns value and presence indicator
- `Delete(key K)` - Removes a key
- `Len() int` - Returns the number of elements
- `Keys() iter.Seq[K]` - Returns an iterator over keys
- `Values() iter.Seq[V]` - Returns an iterator over values
- `All() iter.Seq2[K, V]` - Returns an iterator over key-value pairs
- `Clear()` - Removes all elements

## Thread Safety

**Initialization** is thread-safe, but the types themselves are **not safe for concurrent access** without external synchronization (like Go's built-in `map` type). If you need concurrent access, use external locking or consider `sync.Map` for maps.

## Why?

In Go, maps and channels have nil zero values and must be initialized with `make()` before use. This can be inconvenient when embedding these types in structs. The `zeros` package makes these types work like other zero-value-usable types in Go's standard library.

## Requirements

Go 1.23 or later (for `iter.Seq2` support)

## License

BSD-3-Clause
