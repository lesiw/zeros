# zeros

[![Go Reference](https://pkg.go.dev/badge/lesiw.io/zeros.svg)](https://pkg.go.dev/lesiw.io/zeros)

Zero-valueable wrappers for Go's built-in types and sync patterns.

## Overview

Package `zeros` provides types that are usable at their zero value:

- **`Chan[T]`** and **`Map[K,V]`** auto-initialize on first use, eliminating the need for explicit `make()` calls
- **`OnceValue[T]`** and **`OnceValues[T1, T2]`** provide zero-valueable alternatives to `sync.OnceValue` and `sync.OnceValues`

All types follow the same principle as `bytes.Buffer` and `sync.Mutex`: they work immediately without initialization.

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

### OnceValue

[▶️ Run this example on the Go Playground](https://go.dev/play/p/fUC-ZF_a42W)

Zero-valueable lazy initialization for a single value:

```go
package main

import (
    "fmt"

    "lesiw.io/zeros"
)

type LazyReader struct {
    init zeros.OnceValue[string]
}

func main() {
    var r LazyReader

    // First call executes the function
    data := r.init.Do(func() string {
        fmt.Println("Loading data")
        return "Hello, World!"
    })
    fmt.Println(data)

    // Subsequent calls return cached value
    data = r.init.Do(func() string {
        fmt.Println("This won't print")
        return "Goodbye"
    })
    fmt.Println(data)
}
```

### OnceValues

[▶️ Run this example on the Go Playground](https://go.dev/play/p/hATNJn-XNEM)

Zero-valueable lazy initialization for two values:

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"

    "lesiw.io/zeros"
)

type LazyFile struct {
	once zeros.OnceValues[*os.File, error]
}

func (f *LazyFile) init() (*os.File, error) {
    fmt.Println("Creating temp file")
    return os.CreateTemp("", "example")
}

func (f *LazyFile) Write(p []byte) (int, error) {
    file, err := f.once.Do(f.init)
    if err != nil {
        return 0, fmt.Errorf("failed to open file: %w", err)
    }
    return file.Write(p)
}

func (f *LazyFile) Stat() (os.FileInfo, error) {
    file, err := f.once.Do(f.init)
    if err != nil {
        return nil, fmt.Errorf("failed to stat file: %w", err)
    }
    return file.Stat()
}

func main() {
    var f LazyFile

    info, err := f.Stat()
    if err != nil {
        log.Fatalf("Stat failed: %v", err)
    }
    fmt.Printf("Initial size: %d bytes\n", info.Size())

    if _, err := io.WriteString(&f, "Hello, world!"); err != nil {
        log.Fatalf("Write failed: %v", err)
    }

    info, err = f.Stat()
    if err != nil {
        log.Fatalf("Stat failed: %v", err)
    }
    fmt.Printf("Final size: %d bytes\n", info.Size())
}
```

## Thread Safety

**`OnceValue` and `OnceValues`** are fully thread-safe. The wrapped function is guaranteed to execute exactly once, even with concurrent calls.

**`Chan` and `Map`** have thread-safe initialization, but the types themselves are **not safe for concurrent access** without external synchronization (like Go's built-in `chan` and `map` types). If you need concurrent map access, use external locking or `sync.Map`.

## Why?

In Go, maps and channels have nil zero values and must be initialized with `make()` before use. This can be inconvenient when embedding these types in structs. The `zeros` package makes these types work like other zero-value-usable types in Go's standard library.

## Requirements

Go 1.23 or later (for `iter.Seq2` support)

## License

BSD-3-Clause
