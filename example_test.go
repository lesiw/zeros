package zeros_test

import (
	"fmt"

	"lesiw.io/zeros"
)

func ExampleChan() {
	var ch zeros.Chan[int]

	go func() {
		ch.Send(42)
	}()

	v := ch.Recv()
	fmt.Println(v)
	// Output: 42
}

func ExampleChan_Close() {
	var ch zeros.Chan[string]

	go func() {
		ch.Send("hello")
		ch.Close()
	}()

	for {
		msg, ok := ch.CheckRecv()
		if !ok {
			break
		}
		fmt.Println(msg)
	}
	// Output: hello
}

func ExampleChan_CheckRecv() {
	var ch zeros.Chan[int]

	go func() {
		ch.Send(100)
		ch.Close()
	}()

	if v, ok := ch.CheckRecv(); ok {
		fmt.Println("received:", v)
	}

	if _, ok := ch.CheckRecv(); !ok {
		fmt.Println("channel closed")
	}
	// Output:
	// received: 100
	// channel closed
}

func ExampleChan_TrySend() {
	var ch zeros.Chan[int]

	// TrySend returns false if no receiver is ready
	if !ch.TrySend(42) {
		fmt.Println("no receiver ready")
	}
	// Output: no receiver ready
}

func ExampleChan_TryRecv() {
	var ch zeros.Chan[int]

	// TryRecv returns false if no value is available
	if _, ok := ch.TryRecv(); !ok {
		fmt.Println("no value available")
	}
	// Output: no value available
}

func ExampleMap() {
	var m zeros.Map[string, int]

	m.Set("answer", 42)

	if v, ok := m.CheckGet("answer"); ok {
		fmt.Println(v)
	}
	// Output: 42
}

func ExampleMap_All() {
	var m zeros.Map[string, string]

	m.Set("hello", "world")
	m.Set("foo", "bar")

	for k, v := range m.All() {
		fmt.Println(k, v)
	}
	// Unordered output:
	// hello world
	// foo bar
}

func ExampleMap_Get() {
	var m zeros.Map[string, int]

	m.Set("exists", 100)

	fmt.Println(m.Get("exists"))
	fmt.Println(m.Get("missing"))
	// Output:
	// 100
	// 0
}

func ExampleMap_CheckGet() {
	var m zeros.Map[string, int]

	m.Set("exists", 100)

	if v, ok := m.CheckGet("exists"); ok {
		fmt.Println("found:", v)
	}

	if _, ok := m.CheckGet("missing"); !ok {
		fmt.Println("not found")
	}
	// Output:
	// found: 100
	// not found
}

func ExampleMap_Len() {
	var m zeros.Map[string, int]

	fmt.Println(m.Len())

	m.Set("a", 1)
	m.Set("b", 2)

	fmt.Println(m.Len())
	// Output:
	// 0
	// 2
}

func ExampleMap_Delete() {
	var m zeros.Map[string, int]

	m.Set("key", 100)
	fmt.Println(m.Len())

	m.Delete("key")
	fmt.Println(m.Len())
	// Output:
	// 1
	// 0
}

func ExampleMap_Keys() {
	var m zeros.Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)

	for k := range m.Keys() {
		fmt.Println(k)
	}
	// Unordered output:
	// a
	// b
}

func ExampleMap_Values() {
	var m zeros.Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)

	for v := range m.Values() {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 2
}

func ExampleMap_Clear() {
	var m zeros.Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	fmt.Println(m.Len())

	m.Clear()
	fmt.Println(m.Len())
	// Output:
	// 2
	// 0
}
