package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	oldFlags := log.Flags()
	oldWriter := log.Writer()

	// Set log flags to 0 to omit the timestamp
	log.SetFlags(0)

	// Redirect log output to a pipe
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	cl := &ConsoleLogger{}
	message := "Hello Console Logger"
	cl.Log(message)

	// Close the pipe and restore the original log settings
	w.Close()
	log.SetFlags(oldFlags)
	log.SetOutput(oldWriter)

	// Read the output from the pipe
	out, _ := io.ReadAll(r)
	if string(out) != message+"\n" {
		t.Errorf("Expected %s, but got %s", message, string(out))
	}
}
