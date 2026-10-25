package cmpl

import (
	"encoding/json"
	"os"
)

type write struct {
	data any
}

type cnst struct {
	value int
}

func SaveInstructions(filename string, instructions []any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(instructions)
}
