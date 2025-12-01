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

[▶️ Run this example on the Go Playground](https://go.dev/play/p/h452fJySTjN)

```go
package main

import (
    "fmt"

    "lesiw.io/zeros"
)

func main() {
    var ch zeros.Chan[int]

    go func() {
        ch.C() <- 42 // auto-initializes the channel
    }()

    value := <-ch.C()
    fmt.Println(value)
}
```

Available methods:
- `C() chan T` - Returns the underlying channel (follows `time.Ticker`/`time.Timer` pattern)
- `Close()` - Closes the underlying channel

### Map

[▶️ Run this example on the Go Playground](https://go.dev/play/p/uHckFuMlsyP)

```go
package main

import (
    "fmt"

    "lesiw.io/zeros"
)

func main() {
    var m zeros.Map[string, int]

    m.Set("answer", 42) // auto-initializes the map

    if v, ok := m.Get("answer"); ok {
        fmt.Println(v)
    }

    for k, v := range m.All() {
        fmt.Println(k, v)
    }
}
```

Available methods:
- `Set(key K, value V)` - Sets a key-value pair
- `Get(key K) (V, bool)` - Retrieves a value by key
- `Delete(key K)` - Removes a key
- `Len() int` - Returns the number of elements
- `All() iter.Seq2[K, V]` - Returns an iterator for range-over-func
- `Clear()` - Removes all elements

## Thread Safety

**Initialization** is thread-safe, but the types themselves are **not safe for concurrent access** without external synchronization (like Go's built-in `map` type). If you need concurrent access, use external locking or consider `sync.Map` for maps.

## Why?

In Go, maps and channels have nil zero values and must be initialized with `make()` before use. This can be inconvenient when embedding these types in structs. The `zeros` package makes these types work like other zero-value-usable types in Go's standard library.

## Requirements

Go 1.23 or later (for `iter.Seq2` support)

## License

BSD-3-Clause
