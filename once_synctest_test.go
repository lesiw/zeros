//go:build go1.25

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
			wg.Go(func() {
				if got, want := once.Do(f), 1; got != want {
					t.Errorf("Do(f) = %d, want %d", got, want)
				}
			})
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
			wg.Go(func() {
				v1, v2 := once.Do(f)
				if got, want := v1, 1; got != want {
					t.Errorf("Do(f) v1 = %d, want %d", got, want)
				}
				if got, want := v2, 2; got != want {
					t.Errorf("Do(f) v2 = %d, want %d", got, want)
				}
			})
		}
		wg.Wait()

		if got, want := calls, 1; got != want {
			t.Errorf("calls = %d, want %d", got, want)
		}
	})
}
