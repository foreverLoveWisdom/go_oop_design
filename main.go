package main

import (
	"log"
)

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(message string) {
	log.Println(message)
}

func main() {
	cl := &ConsoleLogger{}
	cl.Log("Hello Console Logger")
}
