package zeros

import "testing"

func TestOnceValuePanic(t *testing.T) {
	var (
		calls int
		once  OnceValue[int]
		f     = func() int {
			calls++
			panic("x")
		}
	)

	for range 10 {
		var p any
		func() {
			defer func() {
				p = recover()
			}()
			once.Do(f)
			t.Fatal("Do(f) did not panic")
		}()
		if got, want := p, "x"; got != want {
			t.Fatalf("panic = %v, want %v", got, want)
		}
	}

	if got, want := calls, 1; got != want {
		t.Errorf("calls = %d, want %d", got, want)
	}
}

func TestOnceValuesPanic(t *testing.T) {
	var (
		calls int
		once  OnceValues[int, int]
		f     = func() (v1, v2 int) {
			calls++
			panic("x")
		}
	)

	for range 10 {
		var p any
		func() {
			defer func() {
				p = recover()
			}()
			once.Do(f)
			t.Fatal("Do(f) did not panic")
		}()
		if got, want := p, "x"; got != want {
			t.Fatalf("panic = %v, want %v", got, want)
		}
	}

	if got, want := calls, 1; got != want {
		t.Errorf("calls = %d, want %d", got, want)
	}
}
