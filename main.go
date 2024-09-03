package main

import (
	"log"
	"os"
)

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(message string) {
	log.Println(message)
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

func (fl *FileLogger) Log(message string) {
	_, err := fl.file.WriteString(message + "\n")
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

func main() {
	cl := &ConsoleLogger{}
	cl.Log("Hello Console Logger")
	fl, err := NewFileLogger("test.log")
	if err != nil {
		log.Println("Error creating file logger:", err)
		return
	}
	defer fl.file.Close()
	defer os.Remove("test.log")
	fl.Log("Hello FileLogger")
	fileContent, err := os.ReadFile("test.log")

	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	log.Println(string(fileContent))
}
