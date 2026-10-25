package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	SaveInstructions()
}

type write struct {
	Data any
}

type cnst struct {
	Value int
}

func SaveInstructions() {
	//constant := cnst{2}
	//writeTest := write{constant}
	data := []any{"CONST", "WRITE"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println(string(jsonData))
}
