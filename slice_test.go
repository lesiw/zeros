package zeros

import (
	"slices"
	"testing"
)

func TestSliceZeroValue(t *testing.T) {
	var s Slice[int]

	s.Append(42)

	if got, want := len(s), 1; got != want {
		t.Errorf("len(s) after Append = %d, want %d", got, want)
	}
	if got, want := s[0], 42; got != want {
		t.Errorf("s[0] = %d, want %d", got, want)
	}
}

func TestSliceAppendVariadic(t *testing.T) {
	var s Slice[int]

	s.Append(1, 2, 3)

	want := []int{1, 2, 3}
	if !slices.Equal(s, want) {
		t.Errorf("s after Append(1, 2, 3) = %v, want %v", s, want)
	}
}

func TestSliceAppendReturnsUpdated(t *testing.T) {
	var s Slice[int]

	got := s.Append(1, 2)

	want := []int{1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("s.Append(1, 2) = %v, want %v", got, want)
	}
}

func TestSliceAppendMultiple(t *testing.T) {
	var s Slice[string]

	s.Append("a")
	s.Append("b", "c")
	s.Append("d")

	want := []string{"a", "b", "c", "d"}
	if !slices.Equal(s, want) {
		t.Errorf("s after appends = %v, want %v", s, want)
	}
}

func TestSliceRangeEmpty(t *testing.T) {
	var s Slice[int]

	var called bool
	for range s {
		called = true
	}
	if called {
		t.Error("range over empty Slice iterated")
	}
}

func TestSliceSlice(t *testing.T) {
	var s Slice[int]
	s.Append(10, 20, 30)

	if got, want := len(s), 3; got != want {
		t.Errorf("len(s) = %d, want %d", got, want)
	}
	if got, want := s[1], 20; got != want {
		t.Errorf("s[1] = %d, want %d", got, want)
	}

	var sum int
	for _, v := range s {
		sum += v
	}
	if got, want := sum, 60; got != want {
		t.Errorf("sum of range = %d, want %d", got, want)
	}

	slices.Reverse(s)
	want := []int{30, 20, 10}
	if !slices.Equal(s, want) {
		t.Errorf("slices.Reverse(s) = %v, want %v", s, want)
	}
}
