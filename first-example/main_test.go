package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

// 1. when we run our program, it's printing it to OS.stdout = standard output
// 2. create a var stdOut -> saving what stdOut is
// 3. create our own stdOut
// 4. spawn out goroutine
// 5. close stdOut
// 5. read from stdOut & convert bytes into strings
// 6. set our create stdOut var to os.Stdout & and test the output
func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("KaLonji", &wg)
	wg.Wait()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	out := string(res)
	os.Stdout = stdOut

	if !strings.Contains(out, "KaLonji") {
		t.Errorf("Expected to find KaLonji, but is not there")
	}
}
