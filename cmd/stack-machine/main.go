package main

import (
	"log"
	"os"

	"interpreter/internal/stackmachine"
)

func main() {
	filePath := "contents.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	instructions, err := stackmachine.LoadProgram(filePath)
	if err != nil {
		log.Fatalf("failed to load program: %v", err)
	}

	in := stackmachine.New(instructions)
	in.Run()
}
