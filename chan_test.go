package zeros

import "testing"

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

func TestChanTryRecvClosed(t *testing.T) {
	var ch Chan[int]

	ch.Close()

	if _, ok := ch.TryRecv(); ok {
		t.Error("Chan[int].TryRecv() on closed channel = true, want false")
	}
}
