package main

import (
	"log"
	"os"

	"interpreter/internal/ast"
)

func main() {
	filePath := "contents.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	program, err := ast.LoadProgram(filePath)
	if err != nil {
		log.Fatalf("failed to load program: %v", err)
	}

	in := ast.New()
	in.Run(program)
}
