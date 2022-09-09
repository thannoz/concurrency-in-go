package main

import "testing"

func Test_UpdateMessage(t *testing.T) {
	msg = "Hello, world"

	wg.Add(2)
	go updateMessage("Goodbye, cruel world!")
	go updateMessage("Yo, whats up?!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}
