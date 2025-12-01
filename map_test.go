package zeros

import (
	"maps"
	"testing"
)

func TestMapZeroValue(t *testing.T) {
	var m Map[string, int]

	m.Set("answer", 42)

	if got, ok := m.Get("answer"); !ok {
		t.Errorf(
			"Map[string, int].Get(%q) returned ok=false, want ok=true",
			"answer",
		)
	} else if want := 42; got != want {
		t.Errorf("Map[string, int].Get(%q) = %d, want %d", "answer", got, want)
	}
}

func TestMapMissing(t *testing.T) {
	var m Map[string, int]

	if _, ok := m.Get("missing"); ok {
		t.Errorf(
			"Map[string, int].Get(%q) returned ok=true, want ok=false",
			"missing",
		)
	}
}

func TestMapDelete(t *testing.T) {
	var m Map[string, int]

	m.Set("key", 100)
	m.Delete("key")

	if _, ok := m.Get("key"); ok {
		t.Errorf(
			"Map[string, int].Get(%q) after Delete(%q) returned "+
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

func TestMapClear(t *testing.T) {
	var m Map[string, int]

	m.Set("a", 1)
	m.Set("b", 2)
	m.Clear()

	if got := m.Len(); got != 0 {
		t.Errorf("Map[string, int].Len() after Clear() = %d, want 0", got)
	}
}
