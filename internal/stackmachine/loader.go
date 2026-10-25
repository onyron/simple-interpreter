package stackmachine

import (
	"encoding/json"
	"os"
)

func LoadProgram(path string) ([]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var instructions []any
	if err := json.Unmarshal(data, &instructions); err != nil {
		return nil, err
	}
	return instructions, nil
}
