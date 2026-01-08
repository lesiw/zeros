package zeros

import (
	"sync"
	"testing"
	"testing/synctest"
)

func TestOnceValue(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var (
			calls int
			once  OnceValue[int]
			wg    sync.WaitGroup
			f     = func() int {
				calls++
				return calls
			}
		)
		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if got, want := once.Do(f), 1; got != want {
					t.Errorf("Do(f) = %d, want %d", got, want)
				}
			}()
		}
		wg.Wait()

		if got, want := calls, 1; got != want {
			t.Errorf("calls = %d, want %d", got, want)
		}
	})
}

func TestOnceValues(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var (
			calls int
			once  OnceValues[int, int]
			wg    sync.WaitGroup
			f     = func() (v1, v2 int) {
				calls++
				return calls, calls + 1
			}
		)

		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				v1, v2 := once.Do(f)
				if got, want := v1, 1; got != want {
					t.Errorf("Do(f) v1 = %d, want %d", got, want)
				}
				if got, want := v2, 2; got != want {
					t.Errorf("Do(f) v2 = %d, want %d", got, want)
				}
			}()
		}
		wg.Wait()

		if got, want := calls, 1; got != want {
			t.Errorf("calls = %d, want %d", got, want)
		}
	})
}

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
