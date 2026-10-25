package main

import (
	"interpreter/internal/cmpl"
	"log"
	"os"
)

func main() {
	filePath := "contents.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	program, err := cmpl.LoadProgram(filePath)
	if err != nil {
		log.Fatalf("failed to load program: %v", err)
	}

	comp := cmpl.New()
	instructions := comp.Compile(program)
	path := "result.json"
	err = cmpl.SaveInstructions(path, instructions)
	if err != nil {
		panic(err)
	}
}
