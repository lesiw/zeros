package zeros

import (
	"testing"
	"testing/synctest"
)

func TestChanZeroValue(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		go func() { ch.C() <- 42 }()

		if got := <-ch.C(); got != 42 {
			t.Errorf("<-Chan[int].C() = %d, want 42", got)
		}
	})
}

func TestChanMultiple(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[string]

		go func() {
			ch.C() <- "hello"
			ch.C() <- "world"
		}()

		if got := <-ch.C(); got != "hello" {
			t.Errorf(
				"<-Chan[string].C() first call = %q, want %q",
				got, "hello",
			)
		}
		if got := <-ch.C(); got != "world" {
			t.Errorf(
				"<-Chan[string].C() second call = %q, want %q",
				got, "world",
			)
		}
	})
}

func TestChanUnderlying(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		go func() { ch.C() <- 99 }()

		if got := <-ch.C(); got != 99 {
			t.Errorf("<-Chan[int].C() = %d, want 99", got)
		}
	})
}

func TestChanClose(t *testing.T) {
	var ch Chan[int]

	ch.Close()

	if _, ok := <-ch.C(); ok {
		t.Error(
			"Chan[int].Close() did not close channel: " +
				"receive succeeded when it should have failed",
		)
	}
}
