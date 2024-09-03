package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestConsoleLogger_Log(t *testing.T) {
	// Save the current log flags and writer
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

func TestFileLogger_Log(t *testing.T) {
	filename := "test.log"
	fl, err := NewFileLogger(filename)
	if err != nil {
		t.Errorf("Error creating file logger: %v", err)
	}
	defer os.Remove(filename)
	message := "Hello FileLogger"
	fl.Log(message)
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	if string(content) != message+"\n" {
		t.Errorf("expected %q but got %q", message+"\n", string(content))
	}
}

func TestNewFileLogger_ErrorHandling(t *testing.T) {
	_, err := NewFileLogger("/root/test.log")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}
