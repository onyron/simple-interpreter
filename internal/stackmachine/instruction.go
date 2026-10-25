package stackmachine

import "fmt"

func (in *Interpreter) handleInstructions() {
	for in.ip = 0; in.ip < len(in.instructions); in.ip++ {
		if instruction, ok := in.instructions[in.ip].(string); ok {
			in.parseInstruction(instruction)
		} else {
			in.parseInstructionMap(in.instructions[in.ip].(map[string]any))
		}
	}
}

func (in *Interpreter) parseInstruction(instruction string) {
	if instruction == "READ" {
		var read int
		_, err := fmt.Scan(&read)
		if err != nil {
			panic("Input error: \"decimal constant\" expected")
		}
		in.stack.Push(read)
	}
	if instruction == "WRITE" {
		fmt.Println(in.stack.Pop())
	}
}

func (in *Interpreter) parseInstructionMap(instruction map[string]any) {
	for key, value := range instruction {
		switch key {
		case "LD":
			in.stack.Push(in.vars[value.(string)])
		case "ST":
			in.vars[value.(string)] = in.stack.Pop()
		case "CONST":
			in.stack.Push(int(value.(float64)))
		case "BINOP":
			in.doBinOp(value.(string))
		case "JMP":
			in.doJmp(value.(string))
		case "JZ":
			in.doJz(value.(string))
		case "JNZ":
			in.doJnz(value.(string))
		}
	}
}

func (in *Interpreter) doJz(label string) {
	if !in.jmpFlag() {
		in.doJmp(label)
	}
}

func (in *Interpreter) doJnz(label string) {
	if in.jmpFlag() {
		in.doJmp(label)
	}
}

func (in *Interpreter) jmpFlag() bool {
	flag := in.stack.Pop()
	if flag == 0 {
		return false
	}
	return true
}

func (in *Interpreter) doJmp(label string) {
	in.ip = in.labels[label]
}

func (in *Interpreter) findLabels() {
	for i := range in.instructions {
		if instruction, ok := in.instructions[i].(map[string]any); ok {
			for key, value := range instruction {
				if key == "LABEL" {
					in.labels[value.(string)] = i
				}
			}
		}
	}
}
