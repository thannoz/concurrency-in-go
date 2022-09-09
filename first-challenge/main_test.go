package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Hello, Saturn!", &wg)
	wg.Wait()

	if !strings.Contains(msg, "Hello, Saturn!") {
		t.Error("Expected to find Hello, Saturn but is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Nonny"
	printMessage()

	_ = w.Close()
	res, _ := io.ReadAll(r)
	out := string(res)
	os.Stdout = stdOut

	if !strings.Contains(out, "Nonny") {
		t.Error("Expected to find Nonny, but is not there")
	}

}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	out := string(res)

	os.Stdout = stdOut

	if !strings.Contains(out, "Hello, universe!") {
		t.Error("Expected to find Hello, universe!, but is not there")
	}

	if !strings.Contains(out, "Hello, cosmos!") {
		t.Error("Expected to find Hello, cosmos!, but is not there")
	}
	if !strings.Contains(out, "Hello, pluto!") {
		t.Error("Expected to find Hello, pluto!, but is not there")
	}
}
