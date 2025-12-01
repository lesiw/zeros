package zeros

import (
	"testing"
	"testing/synctest"
)

func TestChanZeroValue(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		go func() { ch.Send(42) }()

		if got := ch.Recv(); got != 42 {
			t.Errorf("Chan[int].Recv() = %d, want 42", got)
		}
	})
}

func TestChanMultiple(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[string]

		go func() {
			ch.Send("hello")
			ch.Send("world")
		}()

		if got := ch.Recv(); got != "hello" {
			t.Errorf(
				"Chan[string].Recv() first call = %q, want %q",
				got, "hello",
			)
		}
		if got := ch.Recv(); got != "world" {
			t.Errorf(
				"Chan[string].Recv() second call = %q, want %q",
				got, "world",
			)
		}
	})
}

func TestChanSendRecv(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		go func() { ch.Send(99) }()

		if got := ch.Recv(); got != 99 {
			t.Errorf("Chan[int].Recv() = %d, want 99", got)
		}
	})
}

func TestChanClose(t *testing.T) {
	var ch Chan[int]

	ch.Close()

	if _, ok := ch.CheckRecv(); ok {
		t.Error(
			"Chan[int].Close() did not close channel: " +
				"receive succeeded when it should have failed",
		)
	}
}

func TestChanCheckRecv(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		go func() { ch.Send(42) }()

		if got, ok := ch.CheckRecv(); !ok || got != 42 {
			t.Errorf("Chan[int].CheckRecv() = %d, %v, want 42, true", got, ok)
		}
	})
}

func TestChanTrySend(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]
		var received int

		// Start a receiver goroutine
		go func() {
			received = ch.Recv()
		}()

		// Wait for receiver to block
		synctest.Wait()

		// Now try to send - should succeed since receiver is waiting
		if !ch.TrySend(42) {
			t.Error(
				"Chan[int].TrySend(42) with waiting receiver = false, " +
					"want true",
			)
		}

		// Wait for receiver to complete
		synctest.Wait()

		if received != 42 {
			t.Errorf("received = %d, want 42", received)
		}
	})
}

func TestChanTryRecv(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		// Start a sender goroutine
		go func() {
			ch.Send(99)
		}()

		// Wait for sender to block
		synctest.Wait()

		// Now try to receive - should succeed since sender is waiting
		if got, ok := ch.TryRecv(); !ok || got != 99 {
			t.Errorf("Chan[int].TryRecv() = %d, %v, want 99, true", got, ok)
		}
	})
}

func TestChanTryRecvClosed(t *testing.T) {
	var ch Chan[int]

	ch.Close()

	if _, ok := ch.TryRecv(); ok {
		t.Error("Chan[int].TryRecv() on closed channel = true, want false")
	}
}

func TestChanChan(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var ch Chan[int]

		// Get underlying channel
		underlying := ch.Chan()

		// Use it directly
		go func() { underlying <- 42 }()

		if got := <-underlying; got != 42 {
			t.Errorf("underlying channel receive = %d, want 42", got)
		}
	})
}
