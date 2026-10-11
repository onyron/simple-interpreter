package main

import (
	"interpreter/interpreter"
	"log"
)

func main() {
	instructions, err := interpreter.LoadProgram("contents.json")
	if err != nil {
		log.Fatalf("failed to load program: %v", err)
	}

	in := interpreter.New(instructions)
	in.Run()
}
