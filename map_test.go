package zeros

import (
	"maps"
	"testing"
)

func TestMapZeroValue(t *testing.T) {
	var m Map[string, int]

	m.Set("answer", 42)

	if got, ok := m.CheckGet("answer"); !ok {
		t.Errorf(
			"Map[string, int].CheckGet(%q) returned ok=false, want ok=true",
			"answer",
		)
	} else if want := 42; got != want {
		t.Errorf(
			"Map[string, int].CheckGet(%q) = %d, want %d",
			"answer", got, want,
		)
	}
}

func TestMapMissing(t *testing.T) {
	var m Map[string, int]

	if _, ok := m.CheckGet("missing"); ok {
		t.Errorf(
			"Map[string, int].CheckGet(%q) returned ok=true, want ok=false",
			"missing",
		)
	}
}

func TestMapGet(t *testing.T) {
	var m Map[string, int]

	m.Set("exists", 42)
	m.Set("zero", 0)

	if got := m.Get("exists"); got != 42 {
		t.Errorf("Map[string, int].Get(%q) = %d, want 42", "exists", got)
	}

	if got := m.Get("zero"); got != 0 {
		t.Errorf("Map[string, int].Get(%q) = %d, want 0", "zero", got)
	}

	if got := m.Get("missing"); got != 0 {
		t.Errorf("Map[string, int].Get(%q) = %d, want 0", "missing", got)
	}
}

func TestMapDelete(t *testing.T) {
	var m Map[string, int]

	m.Set("key", 100)
	m.Delete("key")

	if _, ok := m.CheckGet("key"); ok {
		t.Errorf(
			"Map[string, int].CheckGet(%q) after Delete(%q) returned "+
				"ok=true, want ok=false",
			"key", "key",
		)
	}
}

func TestMapLen(t *testing.T) {
	var m Map[string, int]

	if got := m.Len(); got != 0 {
		t.Errorf("Map[string, int].Len() = %d, want 0", got)
	}

	m.Set("a", 1)
	m.Set("b", 2)

	if got := m.Len(); got != 2 {
		t.Errorf(
			"Map[string, int].Len() after 2 Set operations = %d, want 2",
			got,
		)
	}
}

func TestMapAll(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	seen := maps.Collect(m.All())

	if len(seen) != 3 {
		t.Errorf("range m.All() saw %d entries, want 3", len(seen))
	}
	want := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, wantV := range want {
		if gotV, ok := seen[k]; !ok {
			t.Errorf("range m.All() did not see key %q", k)
		} else if gotV != wantV {
			t.Errorf(
				"range m.All() for key %q = %d, want %d",
				k, gotV, wantV,
			)
		}
	}
}

func TestMapAllEmpty(t *testing.T) {
	var m Map[string, int]

	var called bool
	for range m.All() {
		called = true
	}

	if called {
		t.Error("range m.All() iterated over empty map")
	}
}

func TestMapAllStop(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	var n int
	for range m.All() {
		n++
		if n >= 2 {
			break
		}
	}

	if n != 2 {
		t.Errorf("range m.All() with break iterated %d times, want 2", n)
	}
}

func TestMapKeys(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	seen := make(map[string]bool)
	for k := range m.Keys() {
		seen[k] = true
	}

	if len(seen) != 3 {
		t.Errorf("range m.Keys() saw %d keys, want 3", len(seen))
	}
	want := []string{"a", "b", "c"}
	for _, k := range want {
		if !seen[k] {
			t.Errorf("range m.Keys() did not see key %q", k)
		}
	}
}

func TestMapValues(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	seen := make(map[int]bool)
	for v := range m.Values() {
		seen[v] = true
	}

	if len(seen) != 3 {
		t.Errorf("range m.Values() saw %d values, want 3", len(seen))
	}
	want := []int{1, 2, 3}
	for _, v := range want {
		if !seen[v] {
			t.Errorf("range m.Values() did not see value %d", v)
		}
	}
}

func TestMapValuesEmpty(t *testing.T) {
	var m Map[string, int]

	var called bool
	for range m.Values() {
		called = true
	}

	if called {
		t.Error("range m.Values() iterated over empty map")
	}
}

func TestMapKeysEmpty(t *testing.T) {
	var m Map[string, int]

	var called bool
	for range m.Keys() {
		called = true
	}

	if called {
		t.Error("range m.Keys() iterated over empty map")
	}
}

func TestMapClear(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Clear()

	if got := m.Len(); got != 0 {
		t.Errorf("Map[string, int].Len() after Clear() = %d, want 0", got)
	}
}

func TestMapMap(t *testing.T) {
	var m Map[string, int]

	m.Set("answer", 42)

	// Get underlying map
	underlying := m.Map()

	// Use it directly
	if got := underlying["answer"]; got != 42 {
		t.Errorf("underlying map[\"answer\"] = %d, want 42", got)
	}

	// Modify through underlying map
	underlying["new"] = 100

	if got := m.Get("new"); got != 100 {
		t.Errorf(
			"m.Get(\"new\") after underlying map modification = %d, want 100",
			got,
		)
	}
}
