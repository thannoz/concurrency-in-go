package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "34320.000000") {
		t.Error("wrong balance")
	}
}
