package main

import (
	"log"
	"os"

	"interpreter/internal/sm"
)

func main() {
	filePath := "contents.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	instructions, err := sm.LoadProgram(filePath)
	if err != nil {
		log.Fatalf("failed to load program: %v", err)
	}

	in := sm.New(instructions)
	in.Run()
}
