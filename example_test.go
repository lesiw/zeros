package zeros_test

import (
	"fmt"

	"lesiw.io/zeros"
)

func ExampleChan() {
	var ch zeros.Chan[int]

	go func() {
		ch.C() <- 42
	}()

	fmt.Println(<-ch.C())
	// Output: 42
}

func ExampleChan_close() {
	var ch zeros.Chan[string]

	go func() {
		ch.C() <- "hello"
		ch.Close()
	}()

	for msg := range ch.C() {
		fmt.Println(msg)
	}
	// Output: hello
}

func ExampleMap() {
	var m zeros.Map[string, int]

	m.Set("answer", 42)

	if v, ok := m.Get("answer"); ok {
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

	if v, ok := m.Get("exists"); ok {
		fmt.Println("found:", v)
	}

	if _, ok := m.Get("missing"); !ok {
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
